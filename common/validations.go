package common

import (
	"fmt"
	"net/mail"
	"net/url"
	"path"
	"regexp"
	"slices"
	"strings"
)

const manifestFile = "funding.json"
const githubHost = "github.com"

var (
	reTag   = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	reID    = regexp.MustCompile(`^[a-z0-9]([[a-z\d-]+)?[a-z0-9]$`)
	rePhone = regexp.MustCompile(`^\+?(\d+-)*\d+$`)
)

// InRange checks whether the given number > min and < max.
func InRange[T int | int32 | int64 | float32 | float64](tag string, num T, min, max T) error {
	if num < min || num > max {
		return fmt.Errorf("`%s` should be of length %v - %v", tag, min, max)
	}

	return nil
}

func InList[S ~[]E, E comparable](tag string, item E, items S) error {
	if !slices.Contains(items, item) {
		if s, ok := any(items).([]string); ok {
			return fmt.Errorf("`%s` should be one of: %s", tag, strings.Join(s, ", "))
		}
		return fmt.Errorf("`%s` has an unknown value", tag)
	}

	return nil
}

func InMap[M ~map[I]I, I comparable](tag, mapName string, item I, mp M) error {
	if _, ok := mp[item]; !ok {
		return fmt.Errorf("`%s` not found in the %s", tag, mapName)
	}

	return nil
}

func MaxItems[T ~[]E, E any](tag string, set T, max int) error {
	if len(set) > max {
		return fmt.Errorf("`%s` can only have max %d elements", tag, max)
	}

	return nil
}

func IsEmail(tag, s string, maxLen int) error {
	if err := InRange(tag, len(s), 3, maxLen); err != nil {
		return err
	}

	em, err := mail.ParseAddress(s)
	if err != nil || em.Address != s {
		return fmt.Errorf("`%s` is not a valid e-mail", tag)
	}

	return nil
}

func IsURL(tag, u string, maxLen int) (*url.URL, error) {
	if err := InRange(tag, len(u), 10, maxLen); err != nil {
		return nil, err
	}

	p, err := url.Parse(u)
	if err != nil || p.Host == "" || (p.Scheme != "https" && p.Scheme != "http") {
		return nil, fmt.Errorf("`%s` is not a valid URL", tag)
	}

	if p.Path != "" {
		p.Path = path.Clean(strings.ReplaceAll(p.Path, "...", ""))
		// p.RawPath = url.PathEscape(p.Path)
	}

	return p, nil
}

// WellKnownState represents the outcome of wellKnown URL validation
type WellKnownState int

const (
	// WellKnownNotRequired indicates the URL is verified without needing a wellKnown URL.
	WellKnownNotRequired WellKnownState = iota

	// WellKnownRequired indicates a wellKnown URL is required but was not provided.
	WellKnownRequired

	// WellKnownValid indicates the provided wellKnown URL is valid.
	WellKnownValid

	// WellKnownInvalid indicates the provided wellKnown URL has errors.
	WellKnownInvalid
)

// isWellKnownRequired checks if a wellKnown URL is required based on manifest and target URLs
func isWellKnownRequired(manifestURL, target *url.URL, mfPath, tgPath string) bool {
	// Different hosts always require wellKnown
	if manifestURL.Host != target.Host {
		return true
	}

	// manifest is in the root of the domain, so all sub-paths are verified.
	//  eg: site.com/funding.json ~= site.com/project
	if mfPath == "/" {
		return false
	}

	// eg: github.com/user/project/blob/main/funding.json  ~= github.com/user/project
	//     /user/project/blob/main/                        ~= /user/project/
	// eg: site.com/project/files/funding.json             ~= site.com/project
	//     /project/files/                                 ~= /project/
	// The manifest path can be in a sub path of the target URL, or vice versa.
	if strings.HasPrefix(tgPath, mfPath) || strings.HasPrefix(mfPath, tgPath) {
		return false
	}

	// Special case for github.com, where github.com/$user/$user is a special case repo where a single
	// funding.json can be hosted, which can be considered valid for all github.com/$user/* repos.
	if manifestURL.Host == githubHost && target.Host == githubHost {
		// Get the username and repo name from the GitHub manifestURL and check if they're the same values.
		// Eg: github.com/user/user -> user == user
		parts := strings.Split(strings.TrimPrefix(mfPath, "/"), "/")
		if len(parts) > 2 && parts[0] == parts[1] {
			// Check if the target URI is a subpath of the /$user URI.
			if strings.HasPrefix(tgPath, fmt.Sprintf("/%s/", parts[0])) {
				return false
			}
		}
	}

	return true
}

// WellKnownURL checks a given targetURL against the given root manifestURL and validates
// the wellKnown URL if provided. Returns the validation result and any error.
func WellKnownURL(tag string, manifestURL *url.URL, target, wellKnown *url.URL, wellKnownURI string) (WellKnownState, error) {
	var (
		// Get the paths and suffix them with "/" for checking using HasPrefix later.
		mfPath = strings.TrimSuffix(strings.TrimSuffix(manifestURL.Path, manifestFile), "/") + "/"
		tgPath = strings.TrimSuffix(target.Path, "/") + "/"
	)

	// Check if wellKnown is required based on host and path matching.
	isRequired := isWellKnownRequired(manifestURL, target, mfPath, tgPath)

	// If wellKnown is not required.
	if !isRequired {
		return WellKnownNotRequired, nil
	}

	// wellKnown is required but not provided.
	if wellKnown == nil {
		return WellKnownRequired, nil
	}

	// Validate the provided wellKnown URL.
	if !strings.HasSuffix(wellKnown.Path, wellKnownURI) {
		return WellKnownInvalid, fmt.Errorf("`%s.wellKnown` should end in %s", tag, wellKnownURI)
	}

	// wellKnown URL should match the main URL.
	if wellKnown.Host != target.Host {
		return WellKnownInvalid, fmt.Errorf("%s.url and `%s.wellKnown` hostnames do not match", tag, tag)
	}

	var (
		wkPath   = strings.TrimSuffix(wellKnown.Path, "/")
		isWKRoot = strings.TrimSuffix(wkPath, wellKnownURI) == ""
	)

	// If wellKnown is at the root of the host, then all sub-paths are acceptable.
	if isWKRoot {
		return WellKnownValid, nil
	}

	// If it's not at the root, then basePath should be a prefix of the well known path.
	// eg:
	// github.com/user ~= github.com/user/project/blob/main/.wellKnown/funding-manifest-urls
	// github.com/user/project ~= github.com/user/project/blob/main/.wellKnown/funding-manifest-urls
	// github.com/use !~= github.com/user/project/blob/main/.wellKnown/funding-manifest-urls
	if tgPath != "/" && !strings.HasPrefix(wkPath, tgPath) {
		return WellKnownInvalid, fmt.Errorf("%s.url and manifest URL host and paths do not match. Expected %s.wellKnown for provenance check at %s://%s%s/*%s", tag, tag, target.Scheme, target.Host, target.Path, wellKnownURI)
	}

	return WellKnownValid, nil
}

func IsRepoURL(tag, u string) error {
	if err := InRange(tag, len(u), 8, 1024); err != nil {
		return err
	}

	p, err := url.Parse(u)
	if err != nil || (p.Scheme != "https" && p.Scheme != "http" && p.Scheme != "git" && p.Scheme != "svn") {
		return fmt.Errorf("%s is not a valid URL", p)
	}

	return nil
}

func IsTag(tag string, val string, min, max int) error {
	if err := InRange(tag, len(val), min, max); err != nil {
		return err
	}

	if !reTag.MatchString(val) {
		return fmt.Errorf("%s should be lowercase alpha-numeric-dashes and length %d - %d", tag, min, max)
	}

	return nil
}

func IsID(tag string, val string, min, max int) error {
	if err := InRange(tag, len(val), min, max); err != nil {
		return err
	}

	err := fmt.Errorf("%s should be lowercase alpha-numeric-dashes and length %d - %d", tag, min, max)

	if !reID.MatchString(val) {
		return err
	}

	if strings.Contains(val, "--") {
		return err
	}

	return nil
}

func IsPhone(tag string, val string) error {
	if len(val) > 32 || !rePhone.MatchString(val) {
		err := fmt.Errorf("%s should only have numbers and optional dashes and max length 32", tag)
		return err
	}

	return nil
}
