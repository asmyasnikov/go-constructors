package http7

import (
	"bytes"
	"fmt"
	"http_request/pkg/xhttp"
	"io"
	"net/http"
	"strings"
)

type option interface {
	apply(*xhttp.Request) error
}
type bodyRequestOption struct {
	body io.Reader
}

func (opt bodyRequestOption) apply(r *xhttp.Request) error {
	switch v := opt.body.(type) {
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
func WithBody(body io.Reader) option {
	return bodyRequestOption{body: body}
}

type methodOption string

func (m methodOption) apply(request *xhttp.Request) error {
	request.Method = string(m)

	return nil
}
func WithMethod(method string) option {
	return methodOption(method)
}

type headersOption struct {
	header string
	values []string
}

func (opt headersOption) apply(request *xhttp.Request) error {
	request.Headers[opt.header] = append(request.Headers[opt.header], opt.values...)

	return nil
}
func WithHeader(header string, values ...string) option {
	return headersOption{
		header: header,
		values: values,
	}
}
func NewRequest(url string, opts ...option) (*xhttp.Request, error) {
	r := &xhttp.Request{
		Method:  http.MethodGet,
		URL:     url,
		Headers: make(map[string][]string),
	}
	for _, opt := range opts {
		if err := opt.apply(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}
