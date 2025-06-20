package common

import (
	"net/url"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestTransformURLs(t *testing.T) {
	f := func(src, target string, match bool) {
		t.Helper()

		u, err := url.Parse(src)
		assert.NoError(t, err)

		u2 := TransformURLOrigin(u)
		if match {
			assert.Equal(t, u2.String(), target)
		} else {
			assert.NotEqual(t, u2.String(), target)
		}
	}

	f("https://github.com/user/project/blob/master/funding.json", "https://github.com/user/project/blob/master/funding.json?raw=true", true)
	f("https://github.com/user/project/raw/master/funding.json", "https://github.com/user/project/raw/master/funding.json?raw=true", true)
	f("https://github.com/user/project/blob/master/funding-manifest-urls", "https://github.com/user/project/blob/master/funding-manifest-urls?raw=true", true)
	f("https://github.com/user/project/blob/main/funding-manifest-urls", "https://github.com/user/project/blob/main/funding-manifest-urls?raw=true", true)
	f("https://github.com/user/project/blob/main/sub/folder/funding-manifest-urls", "https://github.com/user/project/blob/main/sub/folder/funding-manifest-urls?raw=true", true)
	f("https://github.com/user/project/raw/main/sub/folder/funding-manifest-urls", "https://github.com/user/project/raw/main/sub/folder/funding-manifest-urls?raw=true", true)
}

func TestTag(t *testing.T) {
	f := func(tag string, isValid bool) {
		t.Helper()

		err := IsTag("test-tag", tag, 1, 50)
		if !isValid {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}

	// Valid
	f("valid", true)
	f("valid-tag", true)
	f("a", true)
	f("123", true)
	f("tag123", true)
	f("multi-word-tag", true)
	f("a1b2c3", true)
	f("tag-with-numbers-123", true)
	f("a-b-c-d-e-f-g", true)
	f("123-456-789", true)

	// Invalid.
	f("invalid_tag", false)
	f("invalid.tag", false)
	f("invalid tag", false)
	f("invalid@tag", false)
	f("invalid#tag", false)
	f("invalid$tag", false)
	f("invalid%tag", false)
	f("invalid+tag", false)
	f("invalid=tag", false)
	f("invalid/tag", false)
	f("invalid\\tag", false)
	f("CamelCase", false)
	f("UPPERCASE", false)
	f("MixedCase123", false)
	f("-invalid", false)
	f("invalid-", false)
	f("-", false)
	f("", false)
	f("invalid--tag", false)
	f("invalid---tag", false)
}
