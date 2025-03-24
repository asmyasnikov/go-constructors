package http6

import (
	"bytes"
	"fmt"
	"http_request/pkg/xhttp"
	"io"
	"net/http"
	"strings"
)

type option func(r *xhttp.Request) error

func WithGet() option {
	return func(r *xhttp.Request) error {
		r.Method = http.MethodGet

		return nil
	}
}
func WithPost() option {
	return func(r *xhttp.Request) error {
		r.Method = http.MethodGet
		return nil
	}
}
func WithMethod(method string) option {
	return func(r *xhttp.Request) error {
		r.Method = method
		return nil
	}
}
func WithBody(body io.Reader) option {
	return func(r *xhttp.Request) error {
		switch v := body.(type) {
		case *bytes.Buffer:
			r.Body = io.NopCloser(bytes.NewReader(v.Bytes()))
		case *bytes.Reader:
			r.Body = io.NopCloser(v)
		case *strings.Reader:
			r.Body = io.NopCloser(v)
		default:
			return fmt.Errorf("unsupported body type: %T", v)
		}
		return nil
	}
}
func WithHeader(header string, values ...string) option {
	return func(r *xhttp.Request) error {
		r.Headers[header] = append(r.Headers[header], values...)
		return nil
	}
}
func NewRequest(url string, opts ...option) (*xhttp.Request, error) {
	r := &xhttp.Request{
		Method:  http.MethodGet,
		URL:     url,
		Headers: make(map[string][]string),
	}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}
