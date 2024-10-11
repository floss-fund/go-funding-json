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

var (
	reTag   = regexp.MustCompile(`^\p{L}([\p{L}\d-]+)?\p{L}$`)
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
	if err := InRange[int](tag, len(s), 3, maxLen); err != nil {
		return err
	}

	em, err := mail.ParseAddress(s)
	if err != nil || em.Address != s {
		return fmt.Errorf("`%s` is not a valid e-mail", tag)
	}

	return nil
}

func IsURL(tag, u string, maxLen int) (*url.URL, error) {
	if err := InRange[int](tag, len(u), 10, maxLen); err != nil {
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

// WellKnownURL checks a given targetURL against the given root manifestURL. If the target host
// is not the same as the manifest host, or if the target path is not equal to, or is a subpath, of the manifest path,
// then a wellknown URL is expected at the manifest host. The bool indicates whether a wellKnown is required at all or not.
func WellKnownURL(tag string, manifestURL *url.URL, target, wellKnown *url.URL, wellKnownURI string) (bool, error) {
	// If there's a manifestURL, then targetURL should be on the same domain. Otherwise, a well-known URL is mandatory.
	if manifestURL.Host != target.Host && wellKnown == nil {
		return true, fmt.Errorf("%s.url and manifest URL host and paths do not match. Expected %s.wellKnown for provenance check at %s://%s%s/*%s", tag, tag, target.Scheme, target.Host, target.Path, wellKnownURI)
	}

	var (
		// Get the paths are suffix them with "/" for checking using HasPrefix later.
		mfPath = strings.TrimRight(strings.TrimRight(manifestURL.Path, manifestFile), "/") + "/"
		tgPath = strings.TrimRight(target.Path, "/") + "/"
	)

	// If the host + paths match, provenance is verified by default and there is no need for wellKnown.
	if manifestURL.Host == target.Host {
		// manfiest is in the root of the domain, so all sub-paths are verified.
		//  eg: site.com/funding.json ~= site.com/project
		if mfPath == "/" {
			return false, nil
		}

		// eg: github.com/user/project/blob/main/funding.json  ~= github.com/user/project
		//     /user/project/blob/main/                        ~= /user/project/
		// eg: site.com/project/files/funding.json             ~= site.com/project
		//     /project/files/                                 ~= /project/
		// The manifest path can be in a sub path of the target URL, or vice versa.
		if strings.HasPrefix(tgPath, mfPath) || strings.HasPrefix(mfPath, tgPath) {
			return false, nil
		}

		return true, fmt.Errorf("%s.url and manifest URL host and paths do not match. Expected %s.wellKnown for provenance check at %s://%s%s/*%s", tag, tag, target.Scheme, target.Host, target.Path, wellKnownURI)
	}

	if !strings.HasSuffix(wellKnown.Path, wellKnownURI) {
		return true, fmt.Errorf("`%s.wellKnown` should end in %s", tag, wellKnownURI)
	}

	// well-known URL should match the main URL.
	if wellKnown.Host != target.Host {
		return true, fmt.Errorf("%s.url and `%s.wellKnown` hostnames do not match", tag, tag)
	}

	var (
		wkPath   = strings.TrimRight(wellKnown.Path, "/")
		isWKRoot = strings.TrimRight(wkPath, wellKnownURI) == ""
	)

	// If wellKnown is at the root of the host, then all sub-paths are acceptable.
	if isWKRoot {
		return true, nil
	}

	// If it's not at the root, then basePath should be a suffix of the well known path.
	// eg:
	// github.com/user ~= github.com/user/project/blob/main/.well-known/funding-manifest-urls
	// github.com/user/project ~= github.com/user/project/blob/main/.well-known/funding-manifest-urls
	// github.com/use !~= github.com/user/project/blob/main/.well-known/funding-manifest-urls
	if tgPath != "/" && !strings.HasPrefix(wkPath, tgPath) {
		return true, fmt.Errorf("%s.url and manifest URL host and paths do not match. Expected %s.wellKnown for provenance check at %s://%s%s/*%s", tag, tag, target.Scheme, target.Host, target.Path, wellKnownURI)
	}

	return true, nil
}

func IsRepoURL(tag, u string) error {
	if err := InRange[int](tag, len(u), 8, 1024); err != nil {
		return err
	}

	p, err := url.Parse(u)
	if err != nil || (p.Scheme != "https" && p.Scheme != "http" && p.Scheme != "git" && p.Scheme != "svn") {
		return fmt.Errorf("%s is not a valid URL", p)
	}

	return nil
}

func IsTag(tag string, val string, min, max int) error {
	if err := InRange[int](tag, len(val), min, max); err != nil {
		return err
	}

	err := fmt.Errorf("%s should be lowercase alpha-numeric-dashes and length %d - %d", tag, min, max)

	if !reTag.MatchString(val) {
		return err
	}

	if strings.Contains(val, "--") {
		return err
	}

	return nil
}

func IsID(tag string, val string, min, max int) error {
	if err := InRange[int](tag, len(val), min, max); err != nil {
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
	if len(val) > 32 || !reID.MatchString(val) {
		err := fmt.Errorf("%s should only have numbers and optional dashes and max length 32", tag)
		return err
	}

	return nil
}
