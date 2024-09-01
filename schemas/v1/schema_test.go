package v1

import (
	"testing"

	"github.com/floss-fund/go-funding-json/common"
	"github.com/stretchr/testify/assert"
)

const wURI = "/.well-known/funding-manifest-urls"

func TestParseURL(t *testing.T) {
	u := URL{URL: "https://floss.fund"}
	assert.NoError(t, parseURL(&u))
	assert.NotNil(t, u.URLobj)
	assert.Nil(t, u.WellKnownObj)
	assert.Equal(t, u.URLobj.String(), u.URL)

	u = URL{URL: "https://floss.fund", WellKnown: "https://floss.fund/.well-known/test"}
	assert.NoError(t, parseURL(&u))
	assert.NotNil(t, u.URLobj)
	assert.NotNil(t, u.WellKnownObj)
	assert.Equal(t, u.URLobj.String(), u.URL)
	assert.Equal(t, u.WellKnownObj.String(), u.WellKnown)
}

func TestWellKnownURL(t *testing.T) {
	f := func(manifest URL, target URL, errExpected bool) {
		t.Helper()

		assert.NoError(t, parseURL(&target))

		res := common.WellKnownURL("t", manifest.URLobj, target.URLobj, target.WellKnownObj, wURI)
		if errExpected {
			assert.Error(t, res)
		} else {
			assert.NoError(t, res)
		}
	}

	m := URL{URL: "https://floss.fund/funding.json"}
	assert.NoError(t, parseURL(&m))

	f(m, URL{URL: "https://floss.fund"}, false)
	f(m, URL{URL: "https://floss.fund", WellKnown: "https://floss.fund/.well-known/funding-manifest-urls"}, false)
	f(m, URL{URL: "https://floss.fund/sub/project", WellKnown: "https://floss.fund/.well-known/funding-manifest-urls"}, false)
	f(m, URL{URL: "https://floss.fund/project"}, false)
	f(m, URL{URL: "https://floss.fund/user/something/project"}, false)

	m = URL{URL: "https://floss.fund/user/funding.json"}
	assert.NoError(t, parseURL(&m))
	f(m, URL{URL: "https://floss.fund/project"}, true)
	f(m, URL{URL: "https://floss.fund/user2/project"}, true)
	f(m, URL{URL: "https://floss.fund/user/project"}, false)
	f(m, URL{URL: "https://floss.fund/user/project/subproject"}, false)

	m = URL{URL: "https://community.org/funding.json"}
	assert.NoError(t, parseURL(&m))
	f(m, URL{URL: "https://project.net"}, true)
	f(m, URL{URL: "https://project.community.org"}, true)

	m = URL{URL: "https://github.com/user/project/blob/main/funding.json"}
	assert.NoError(t, parseURL(&m))
	f(m, URL{URL: "https://floss.fund/project"}, true)
	f(m, URL{URL: "https://github.com/user/project2"}, true)
	f(m, URL{URL: "https://github.com/user2/project"}, true)
	f(m, URL{URL: "https://github.com/user/project"}, false)
}
