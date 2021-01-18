package httpz

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/absurdlab/x/v2/tracer"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Get is NewRequest with http GET method.
func Get(ctx context.Context, url string, options ...RequestOpt) (*http.Request, error) {
	req, err := NewRequest(ctx, http.MethodGet, url, options...)
	if err != nil {
		return nil, tracer.Touch(err)
	}
	return req, nil
}

// Post is NewRequest with http POST method.
func Post(ctx context.Context, url string, options ...RequestOpt) (*http.Request, error) {
	req, err := NewRequest(ctx, http.MethodPost, url, options...)
	if err != nil {
		return nil, tracer.Touch(err)
	}
	return req, nil
}

// Put is NewRequest with http PUT method.
func Put(ctx context.Context, url string, options ...RequestOpt) (*http.Request, error) {
	req, err := NewRequest(ctx, http.MethodPut, url, options...)
	if err != nil {
		return nil, tracer.Touch(err)
	}
	return req, nil
}

// Delete is NewRequest with http DELETE method.
func Delete(ctx context.Context, url string, options ...RequestOpt) (*http.Request, error) {
	req, err := NewRequest(ctx, http.MethodDelete, url, options...)
	if err != nil {
		return nil, tracer.Touch(err)
	}
	return req, nil
}

// NewRequest creates a new http.Request with the given method and url, and applies a series of RequestOpt to it to
// further configure the http.Request.
func NewRequest(ctx context.Context, method string, url string, options ...RequestOpt) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, tracer.Touch(err)
	}

	for _, opt := range options {
		if err = opt(req); err != nil {
			return nil, tracer.Touch(err)
		}
	}

	return req, nil
}

// RequestOpt modifies the http.Request.
type RequestOpt func(req *http.Request) error

var (
	// Body reads the io.Reader and set it as body on the http.Request.
	Body = func(body io.Reader) RequestOpt {
		return func(req *http.Request) error {
			rc, ok := body.(io.ReadCloser)
			if !ok && body != nil {
				rc = ioutil.NopCloser(body)
			}

			req.Body = rc

			if body != nil {
				switch v := body.(type) {
				case *bytes.Buffer:
					req.ContentLength = int64(v.Len())
					buf := v.Bytes()
					req.GetBody = func() (io.ReadCloser, error) {
						r := bytes.NewReader(buf)
						return ioutil.NopCloser(r), nil
					}
				case *bytes.Reader:
					req.ContentLength = int64(v.Len())
					snapshot := *v
					req.GetBody = func() (io.ReadCloser, error) {
						r := snapshot
						return ioutil.NopCloser(&r), nil
					}
				case *strings.Reader:
					req.ContentLength = int64(v.Len())
					snapshot := *v
					req.GetBody = func() (io.ReadCloser, error) {
						r := snapshot
						return ioutil.NopCloser(&r), nil
					}
				default:
					// This is where we'd set it to -1 (at least
					// if body != NoBody) to mean unknown, but
					// that broke people during the Go 1.8 testing
					// period. People depend on it being 0 I
					// guess. Maybe retry later. See Issue 18117.
				}
				// For client requests, Request.ContentLength of 0
				// means either actually 0, or unknown. The only way
				// to explicitly say that the ContentLength is zero is
				// to set the Body to nil. But turns out too much code
				// depends on NewRequest returning a non-nil Body,
				// so we use a well-known ReadCloser variable instead
				// and have the http package also treat that sentinel
				// variable to mean explicitly zero.
				if req.GetBody != nil && req.ContentLength == 0 {
					req.Body = http.NoBody
					req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
				}
			}

			return nil
		}
	}

	// JSON serializes given body as json and sets application/json header.
	JSON = func(body interface{}) RequestOpt {
		return func(req *http.Request) error {
			req.Header.Set("Content-Type", "application/json")

			raw, err := json.Marshal(body)
			if err != nil {
				return tracer.Touch(err)
			}

			return Body(bytes.NewReader(raw))(req)
		}
	}

	// URLForm serializes given values as url encoded form body and sets application/x-www-form-urlencoded header.
	URLForm = func(v url.Values) RequestOpt {
		return func(req *http.Request) error {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			return Body(strings.NewReader(v.Encode()))(req)
		}
	}

	// BasicAuth sets the basic authorization header.
	BasicAuth = func(username, password string) RequestOpt {
		return func(req *http.Request) error {
			h := fmt.Sprintf("%s:%s", username, password)
			h = base64.StdEncoding.EncodeToString([]byte(h))
			req.Header.Set("Authorization", "Basic "+h)
			return nil
		}
	}

	// BearerAuth sets the bearer token as authorization header.
	BearerAuth = func(token string) RequestOpt {
		return func(req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			return nil
		}
	}
)
