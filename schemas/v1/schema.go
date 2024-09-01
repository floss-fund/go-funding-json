package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/floss-fund/go-funding-json/common"
	"golang.org/x/mod/semver"
)

const (
	// MajorVersion is the current major version of the schema definition (v1, v2 etc.)
	MajorVersion = "v1"

	// CurrentVersion is the exact current version of the schema with minor/patch changes.
	// It should be backwards compatible be MajorVersion.
	CurrentVersion = "v1.0.0"

	maxUrlLen = 1024
)

// Schema represents the schema+parser+validator for a particular version.
type Schema struct {
	opt *Opt
	hc  *common.HTTPClient
}

type Opt struct {
	// Map of SPDX ID: License name.
	Licenses map[string]string

	// Map of programming language names.
	ProgrammingLanguages map[string]string

	// Map of curency code and names.
	Currencies map[string]string

	WellKnownURI string
}

// New returns a new instance of Schema.
func New(opt *Opt, hOpt common.HTTPOpt, l *log.Logger) *Schema {
	hc := common.NewHTTPClient(hOpt, l)

	return &Schema{
		opt: opt,
		hc:  hc,
	}
}

// ParseManifest parses a given JSON body, validates it, and returns the manifest.
func (s *Schema) ParseManifest(b []byte, manifestURL string, checkProvenance bool) (Manifest, error) {
	var m Manifest
	if err := m.UnmarshalJSON(b); err != nil {
		return m, fmt.Errorf("error parsing JSON body: %v", err)
	}

	// Validate the manifest's schema.
	m.Body = json.RawMessage(b)

	m.URL = URL{URL: manifestURL}
	if err := parseURL(&m.URL); err != nil {
		return m, err
	}

	if err := parseURL(&m.Entity.WebpageURL); err != nil {
		return m, err
	}

	// Parse various URL strings to url.URL obijects.
	for n := 0; n < len(m.Projects); n++ {
		// Project webpage.
		if err := parseURL(&m.Projects[n].WebpageURL); err != nil {
			return m, err
		}

		// Project repository.
		if err := parseURL(&m.Projects[n].RepositoryUrl); err != nil {
			return m, err
		}
	}

	if v, err := s.Validate(m); err != nil {
		return v, err
	} else {
		m = v
	}

	// Establish the provenance of all URLs mentioned in the manifest.
	if checkProvenance {
		if err := s.CheckProvenance(m.Entity.WebpageURL, m.URL); err != nil {
			return m, err
		}

		for _, o := range m.Projects {
			o := o
			if err := s.CheckProvenance(o.WebpageURL, m.URL); err != nil {
				return m, err
			}
			if err := s.CheckProvenance(o.RepositoryUrl, m.URL); err != nil {
				return m, err
			}
		}
	}

	return m, nil
}

// Validate validates a given manifest against its schema.
func (s *Schema) Validate(m Manifest) (Manifest, error) {
	if semver.Major(m.Version) != MajorVersion {
		return m, fmt.Errorf("major version should be %s (current version is %s)", MajorVersion, CurrentVersion)
	}

	mURL, err := common.IsURL("manifest URL", m.URL.URL, maxUrlLen)
	if err != nil {
		return m, err
	}

	// Entity.
	if m.Entity, err = s.ValidateEntity(m.Entity, mURL); err != nil {
		return m, err
	}

	// Projects.
	for n, o := range m.Projects {
		if o, err = s.ValidateProject(o, n, mURL); err != nil {
			return m, err
		}
		m.Projects[n] = o
	}

	// Funding channels.
	chIDs := make(map[string]struct{})
	for n, o := range m.Funding.Channels {
		if o, err = s.ValidateChannel(o, n); err != nil {
			return m, err
		}

		m.Funding.Channels[n] = o
		chIDs[o.ID] = struct{}{}
	}

	// Funding plans.
	if err := common.InRange[int]("plans", len(m.Funding.Plans), 1, 30); err != nil {
		return m, err
	}
	for n, o := range m.Funding.Plans {
		if o, err = s.ValidatePlan(o, n, chIDs); err != nil {
			return m, err
		}
		m.Funding.Plans[n] = o
	}

	// History.
	if err := common.InRange[int]("history", len(m.Funding.Plans), 0, 50); err != nil {
		return m, err
	}
	for n, o := range m.Funding.History {
		if o, err = s.ValidateHistory(o, n); err != nil {
			return m, err
		}
		m.Funding.History[n] = o
	}

	return m, nil
}

func (s *Schema) ValidateEntity(o Entity, manifest *url.URL) (Entity, error) {
	if err := common.InList("entity.type", o.Type, EntityTypes); err != nil {
		return o, err
	}

	if err := common.InList("entity.role", o.Role, EntityRoles); err != nil {
		return o, err
	}

	if err := common.InRange[int]("entity.name", len(o.Name), 2, 128); err != nil {
		return o, err
	}

	if err := common.IsEmail("entity.email", o.Email, 128); err != nil {
		return o, err
	}

	if err := common.InRange[int]("entity.telephone", len(o.Telephone), 0, 24); err != nil {
		return o, err
	}

	if err := common.WellKnownURL("entity.webpageUrl", manifest, o.WebpageURL.URLobj, o.WebpageURL.WellKnownObj, s.opt.WellKnownURI); err != nil {
		return o, err
	}

	return o, nil
}

func (s *Schema) ValidateProject(o Project, n int, manifest *url.URL) (Project, error) {
	if err := common.InRange[int](fmt.Sprintf("projects[%d].name", n), len(o.Name), 1, 256); err != nil {
		return o, err
	}

	if err := common.InRange[int](fmt.Sprintf("projects[%d].description", n), len(o.Description), 5, maxUrlLen); err != nil {
		return o, err
	}

	if err := common.WellKnownURL(fmt.Sprintf("projects[%d].webpageUrl", n), manifest, o.WebpageURL.URLobj, o.WebpageURL.WellKnownObj, s.opt.WellKnownURI); err != nil {
		return o, err
	}

	if err := common.WellKnownURL(fmt.Sprintf("projects[%d].repositoryUrl", n), manifest, o.RepositoryUrl.URLobj, o.RepositoryUrl.WellKnownObj, s.opt.WellKnownURI); err != nil {
		return o, err
	}

	// License.
	licenseTag := fmt.Sprintf("projects[%d].license", n)
	if err := common.InRange[int](licenseTag, len(o.License), 2, 64); err != nil {
		return o, err
	}
	if strings.HasPrefix(o.License, "spdx:") {
		if err := common.InMap(licenseTag, "spdx license list", strings.TrimPrefix(o.License, "spdx:"), s.opt.Licenses); err != nil {
			return o, err
		}
	}

	// Frameworks.
	if err := common.InRange[int](fmt.Sprintf("projects[%d].frameworks", n), len(o.Frameworks), 0, 5); err != nil {
		return o, err
	}
	for i, f := range o.Frameworks {
		fTag := fmt.Sprintf("projects[%d].frameworks[%d]", n, i)
		if err := common.InRange[int](fTag, len(f), 2, 64); err != nil {
			return o, err
		}

		if strings.HasPrefix(f, "lang:") {
			if err := common.InMap(fTag, "default programming language list", strings.TrimPrefix(f, "lang:"), s.opt.ProgrammingLanguages); err != nil {
				return o, err
			}
		}
	}

	// Tags.
	if err := common.InRange[int](fmt.Sprintf("projects[%d].tags", n), len(o.Tags), 1, 10); err != nil {
		return o, err
	}
	for i, t := range o.Tags {
		if err := common.IsTag(fmt.Sprintf("projects[%d].tags[%d]", n, i), t, 2, 32); err != nil {
			return o, err
		}
	}

	return o, nil
}

func (s *Schema) ValidateChannel(o Channel, n int) (Channel, error) {
	if err := common.IsID(fmt.Sprintf("channels[%d].id", n), o.ID, 3, 32); err != nil {
		return o, err
	}

	if err := common.InList(fmt.Sprintf("channels[%d].type", n), o.Type, ChannelTypes); err != nil {
		return o, err
	}

	if err := common.InRange[int](fmt.Sprintf("channels[%d].address", n), len(o.Address), 0, 128); err != nil {
		return o, err
	}

	if err := common.InRange[int](fmt.Sprintf("channels[%d].description", n), len(o.Description), 0, maxUrlLen); err != nil {
		return o, err
	}

	return o, nil
}

func (s *Schema) ValidatePlan(o Plan, n int, channelIDs map[string]struct{}) (Plan, error) {
	if err := common.IsID(fmt.Sprintf("plans[%d].id", n), o.ID, 3, 32); err != nil {
		return o, err
	}

	if err := common.InList(fmt.Sprintf("plans[%d].status", n), o.Status, PlanStatuses); err != nil {
		return o, err
	}

	if err := common.InRange[int](fmt.Sprintf("plans[%d].name", n), len(o.Name), 3, 128); err != nil {
		return o, err
	}

	if err := common.InRange[int](fmt.Sprintf("plans[%d].description", n), len(o.Description), 0, maxUrlLen); err != nil {
		return o, err
	}

	if err := common.InRange[float64](fmt.Sprintf("plans[%d].amount", n), o.Amount, 0, 1000000000); err != nil {
		return o, err
	}

	if err := common.InMap(fmt.Sprintf("plans[%d].currency", n), "currencies list", o.Currency, s.opt.Currencies); err != nil {
		return o, err
	}

	if err := common.InList(fmt.Sprintf("plans[%d].frequency", n), o.Frequency, PlanFrequencies); err != nil {
		return o, err
	}

	for _, ch := range o.Channels {
		if _, ok := channelIDs[ch]; !ok {
			return o, fmt.Errorf("unknown channel id in plans[%d].frequency", n)
		}
	}

	return o, nil
}

func (s *Schema) ValidateHistory(o HistoryItem, n int) (HistoryItem, error) {
	if err := common.InRange[int](fmt.Sprintf("history[%d].year", n), o.Year, 1970, 2075); err != nil {
		return o, err
	}

	if err := common.InRange[float64](fmt.Sprintf("plans[%d].income", n), o.Income, 0, 1000000000); err != nil {
		return o, err
	}

	if err := common.InRange[float64](fmt.Sprintf("plans[%d].expenses", n), o.Expenses, 0, 1000000000); err != nil {
		return o, err
	}

	if err := common.InRange[int](fmt.Sprintf("projects[%d].description", n), len(o.Description), 0, maxUrlLen); err != nil {
		return o, err
	}

	return o, nil
}

// CheckProvenance fetches the .well-known URL list for the given u and checks
// wehther the manifestURL is present in it, establishing its provenance.
func (s *Schema) CheckProvenance(u URL, manifest URL) error {
	if u.WellKnown == "" {
		return nil
	}

	body, err := s.hc.Get(u.WellKnownObj)
	if err != nil {
		return err
	}

	mStr := manifest.URLobj.String()
	ub := []byte(mStr)
	for n, b := range bytes.Split(body, []byte("\n")) {
		if bytes.Equal(ub, b) {
			return nil
		}

		if n > 100 {
			return errors.New("too many lines in the .well-known list")
		}
	}

	return fmt.Errorf("manifest URL %s was not found in the .well-known list", mStr)
}

func parseURL(u *URL) error {
	{
		p, err := common.IsURL("", u.URL, maxUrlLen)
		if err != nil {
			return err
		}
		u.URLobj = p
	}

	if u.WellKnown != "" {
		p, err := common.IsURL("", u.WellKnown, maxUrlLen)
		if err != nil {
			return err
		}
		u.WellKnownObj = p
	}

	return nil
}
