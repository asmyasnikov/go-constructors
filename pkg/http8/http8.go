package http8

import (
	"bytes"
	"fmt"
	"http_request/pkg/xhttp"
	"io"
	"net/http"
	"strings"
)

type requestBuilder struct {
	r   xhttp.Request
	err error
}

func (b requestBuilder) WithMethod(method string) requestBuilder {
	if b.err != nil {
		return b
	}

	b.r.Method = method

	return b
}
func (b requestBuilder) WithHeader(header string, values ...string) requestBuilder {
	if b.err != nil {
		return b
	}

	b.r.Headers[header] = append(b.r.Headers[header], values...)

	return b
}
func (b requestBuilder) WithBody(body io.Reader) requestBuilder {
	switch v := body.(type) {
	case *bytes.Buffer:
		b.r.Body = io.NopCloser(bytes.NewReader(v.Bytes()))
	case *bytes.Reader:
		b.r.Body = io.NopCloser(v)
	case *strings.Reader:
		b.r.Body = io.NopCloser(v)
	default:
		b.err = fmt.Errorf("unsupported body type: %T", v)
	}
	return b
}
func (b requestBuilder) Build() (*xhttp.Request, error) {
	return &b.r, b.err
}
func NewRequest(url string) requestBuilder {
	return requestBuilder{r: xhttp.Request{
		Method:  http.MethodGet,
		URL:     url,
		Headers: make(map[string][]string),
	}, err: nil}
}
