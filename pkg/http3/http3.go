package http3

import (
	"fmt"
	"http_request/pkg/xhttp"
	"io"
	"net/http"
)

func NewRequest(url string, args ...any) (*xhttp.Request, error) {
	r := &xhttp.Request{
		Method: http.MethodGet,
		URL:    url,
	}

	for _, arg := range args {
		switch v := arg.(type) {
		case http.Header:
			r.Headers = v
		case io.Reader:
			r.Body = io.NopCloser(v)
		case string:
			r.Method = v
		default:
			return nil, fmt.Errorf("unsupported arg type: %T", v)
		}
	}

	return r, nil
}
