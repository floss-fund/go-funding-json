package common

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type HTTPOpt struct {
	UserAgent           string        `json:"useragent"`
	MaxHostConns        int           `json:"max_host_conns"`
	ReqTimeout          time.Duration `json:"req_timeout"`
	Retries             int           `json:"retries"`
	RetryWait           time.Duration `json:"retry_wait"`
	MaxBytes            int64         `json:"max_bytes"`
	SkipRateLimitedHost bool          `json:"skip_ratelimited_host"`
}

type HTTPClient struct {
	rateLimited map[string]struct{}
	headers     http.Header

	opt    HTTPOpt
	client *http.Client
	log    *log.Logger
}

var (
	ErrRatelimited = errors.New("host rate limited the request")
)

// NewHTTPClient returns an instance of an HTTP client that is used for fetching
// manifests and .well-known URLs for checking provenance.
func NewHTTPClient(o HTTPOpt, l *log.Logger) *HTTPClient {
	h := http.Header{}
	h.Set("User-Agent", o.UserAgent)

	return &HTTPClient{
		headers: h,
		opt:     o,
		client: &http.Client{
			Timeout: o.ReqTimeout,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   o.MaxHostConns,
				MaxConnsPerHost:       o.MaxHostConns,
				ResponseHeaderTimeout: o.ReqTimeout,
				IdleConnTimeout:       o.ReqTimeout,
			},
		},
		log: l}
}

// Get fetches a given URL with error retries.
func (h *HTTPClient) Get(u *url.URL) ([]byte, error) {
	var (
		body       []byte
		err        error
		statusCode int
		retry      bool
	)

	// Host is disabled due to rate limiting.
	if _, ok := h.rateLimited[u.Host]; ok {
		return nil, ErrRatelimited
	}

	// Retry N times.
	for n := 0; n < h.opt.Retries; n++ {
		body, _, retry, statusCode, err = h.DoReq(http.MethodGet, u.String(), nil, h.headers)
		if err == nil || !retry {
			break
		}

		// If the host sent a 429, don't send any more requests.
		if h.opt.SkipRateLimitedHost && statusCode == http.StatusTooManyRequests {
			h.rateLimited[u.Host] = struct{}{}
			return nil, ErrRatelimited
		}

		if h.opt.Retries > 1 {
			time.Sleep(h.opt.RetryWait)
		}
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Head fetches the metadata (HEAD) request of a given URL.
func (h *HTTPClient) Head(u *url.URL) (http.Header, error) {
	var (
		hdr        http.Header
		err        error
		statusCode int
		retry      bool
	)

	// Host is disabled due to rate limiting.
	if _, ok := h.rateLimited[u.Host]; ok {
		return nil, ErrRatelimited
	}

	// Retry N times.
	for n := 0; n < h.opt.Retries; n++ {
		_, hdr, retry, statusCode, err = h.DoReq(http.MethodHead, u.String(), nil, h.headers)
		if err == nil || !retry {
			break
		}

		// If the host sent a 429, don't send any more requests.
		if h.opt.SkipRateLimitedHost && statusCode == http.StatusTooManyRequests {
			h.rateLimited[u.Host] = struct{}{}
			return nil, ErrRatelimited
		}

		if h.opt.Retries > 1 {
			time.Sleep(h.opt.RetryWait)
		}
	}
	if err != nil {
		return nil, err
	}

	return hdr, nil
}

// DoReq executes an HTTP request. The bool indicates whether it's a retriable error.
func (h *HTTPClient) DoReq(method, rURL string, reqBody []byte, headers http.Header) (respBody []byte, hdr http.Header, retry bool, statusCode int, retErr error) {
	var (
		err      error
		postBody io.Reader
	)

	defer func() {
		msg := "OK"
		if retErr != nil {
			msg = retErr.Error()
		} else if statusCode != http.StatusOK {
			msg = "FAILED"
		}

		h.log.Printf("%s %s -> %d: %v", method, rURL, statusCode, msg)
	}()

	// Encode POST / PUT params.
	if method == http.MethodPost || method == http.MethodPut {
		postBody = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, rURL, postBody)
	if err != nil {
		return nil, nil, true, 0, err
	}

	if headers != nil {
		req.Header = headers
	} else {
		req.Header = http.Header{}
	}

	// If a content-type isn't set, set the default one.
	if req.Header.Get("Content-Type") == "" {
		if method == http.MethodPost || method == http.MethodPut {
			req.Header.Add("Content-Type", "application/json")
		}
	}

	// If the request method is GET or DELETE, add the params as QueryString.
	if method == http.MethodGet || method == http.MethodDelete {
		req.URL.RawQuery = string(reqBody)
	}

	r, err := h.client.Do(req)
	if err != nil {
		return nil, nil, true, 0, err
	}

	defer func() {
		// Drain and close the body to let the Transport reuse the connection
		io.Copy(io.Discard, r.Body)
		r.Body.Close()
	}()

	body, err := io.ReadAll(io.LimitReader(r.Body, h.opt.MaxBytes))
	if err != nil {
		return nil, nil, true, http.StatusOK, err
	}

	if r.StatusCode > 299 {
		return body, r.Header, false, r.StatusCode, fmt.Errorf("error: %s returned %d", rURL, r.StatusCode)
	}

	return body, r.Header, false, http.StatusOK, nil
}
