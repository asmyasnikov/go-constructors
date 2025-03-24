package http5

import (
	"http_request/pkg/xhttp"
	"net/http"
	"strings"
)

type Settings struct {
	Method  *string
	Proto   *string
	Body    *string
	Headers map[string][]string
}

func NewRequest(url string, settings *Settings) *xhttp.Request {
	r := &xhttp.Request{
		Method: http.MethodGet,
		URL:    url,
	}
	if settings == nil {
		return r
	}
	if settings.Method != nil {
		r.Method = *settings.Method
	}
	if settings.Body != nil {
		r.Body = strings.NewReader(*settings.Body)
	}
	if settings.Headers != nil {
		r.Headers = settings.Headers
	}
	return r
}
