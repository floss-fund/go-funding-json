package v1

import (
	"testing"

	"github.com/floss-fund/go-funding-json/common"
	"github.com/stretchr/testify/assert"
)

const wURI = "/.well-known/funding-manifest-urls"

func TestParseURL(t *testing.T) {
	u := URL{URL: "https://floss.fund"}
	assert.NoError(t, parseURL("", &u))
	assert.NotNil(t, u.URLobj)
	assert.Nil(t, u.WellKnownObj)
	assert.Equal(t, u.URLobj.String(), u.URL)

	u = URL{URL: "https://floss.fund", WellKnown: "https://floss.fund/.well-known/test"}
	assert.NoError(t, parseURL("", &u))
	assert.NotNil(t, u.URLobj)
	assert.NotNil(t, u.WellKnownObj)
	assert.Equal(t, u.URLobj.String(), u.URL)
	assert.Equal(t, u.WellKnownObj.String(), u.WellKnown)
}

func TestWellKnownURL(t *testing.T) {
	f := func(manifest, target URL, expectedState common.WellKnownState, expectError bool) {
		t.Helper()

		assert.NoError(t, parseURL("", &target))

		result, err := common.WellKnownURL("t", manifest.URLobj, target.URLobj, target.WellKnownObj, wURI)
		if expectError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
		assert.Equal(t, expectedState, result)
	}

	// Same host, manifest at root: wellKnown not required.
	m := URL{URL: "https://floss.fund/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://floss.fund"}, common.WellKnownNotRequired, false)
	f(m, URL{URL: "https://floss.fund/project"}, common.WellKnownNotRequired, false)
	f(m, URL{URL: "https://floss.fund/user/something/project"}, common.WellKnownNotRequired, false)

	// Same host with path matching: wellKnown not required.
	m = URL{URL: "https://floss.fund/user/project/blob/main/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://floss.fund/user/project"}, common.WellKnownNotRequired, false)

	m = URL{URL: "https://floss.fund/user/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://floss.fund/user/project"}, common.WellKnownNotRequired, false)
	f(m, URL{URL: "https://floss.fund/user/project/subproject"}, common.WellKnownNotRequired, false)

	// Same host with non-matching paths: wellKnown required.
	m = URL{URL: "https://floss.fund/user/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://floss.fund/project"}, common.WellKnownRequired, false)
	f(m, URL{URL: "https://floss.fund/user2/project"}, common.WellKnownRequired, false)

	// Different hosts: wellKnown required.
	m = URL{URL: "https://example.com/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://www.example.com"}, common.WellKnownRequired, false)
	f(m, URL{URL: "https://project.net"}, common.WellKnownRequired, false)
	f(m, URL{URL: "https://project.example.com"}, common.WellKnownRequired, false)

	m = URL{URL: "https://github.com/user/project/blob/main/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://floss.fund/project"}, common.WellKnownRequired, false)

	// GitHub special case: wellKnown/$user/$user.
	m = URL{URL: "https://github.com/user/user/blob/main/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://github.com/user/project"}, common.WellKnownNotRequired, false)

	m = URL{URL: "https://github.com/johndoe/johndoe/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://github.com/johndoe/awesome-project"}, common.WellKnownNotRequired, false)

	m = URL{URL: "https://github.com/user/user/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://github.com/otheruser/project"}, common.WellKnownRequired, false)

	// Well-known URL validation when provided.
	m = URL{URL: "https://floss.fund/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://floss.fund", WellKnown: "https://floss.fund/.well-known/funding-manifest-urls"}, common.WellKnownNotRequired, false)
	f(m, URL{URL: "https://floss.fund/sub/project", WellKnown: "https://floss.fund/.well-known/funding-manifest-urls"}, common.WellKnownNotRequired, false)

	m = URL{URL: "https://floss.fund/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://example.com", WellKnown: "https://example.com/.well-known/funding-manifest-urls"}, common.WellKnownValid, false)
	f(m, URL{URL: "https://example.com/sub/project", WellKnown: "https://example.com/.well-known/funding-manifest-urls"}, common.WellKnownValid, false)

	// Invalid well-known URLs.
	m = URL{URL: "https://floss.fund/user/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://other.com/project", WellKnown: "https://other.com/.well-known/wrong-suffix"}, common.WellKnownInvalid, true)
	f(m, URL{URL: "https://project.net/path", WellKnown: "https://other.net/.well-known/funding-manifest-urls"}, common.WellKnownInvalid, true)
	f(m, URL{URL: "https://other.com/project/path", WellKnown: "https://other.com/different/.well-known/funding-manifest-urls"}, common.WellKnownInvalid, true)

	// Edge cases.
	m = URL{URL: "https://github.com/user/project/blob/main/funding.json"}
	assert.NoError(t, parseURL("", &m))
	f(m, URL{URL: "https://github.com/user/project/../../test"}, common.WellKnownRequired, false)
	f(m, URL{URL: "https://github.com/user2/project"}, common.WellKnownRequired, false)
	f(m, URL{URL: "https://github.com/user/project2"}, common.WellKnownRequired, false)
}
