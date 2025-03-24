package http4

import (
	"http_request/pkg/xhttp"
	"net/http"
	"strings"
)

type Settings struct {
	Method  string // "GET"
	Body    string
	Headers map[string][]string // nil
}

func NewRequest(url string, settings *Settings) *xhttp.Request {
	r := &xhttp.Request{
		Method: http.MethodGet,
		URL:    url,
	}
	if settings == nil {
		return r
	}
	if settings.Method != "" {
		r.Method = settings.Method
	}
	if settings.Body != "" {
		r.Body = strings.NewReader(settings.Body)
	}
	if settings.Headers != nil {
		r.Headers = settings.Headers
	}
	return r
}
