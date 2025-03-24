package http1

import (
	"io"
	"net/http"

	"http_request/pkg/xhttp"
)

func NewRequest(
	url string,
) *xhttp.Request {
	return &xhttp.Request{
		Method: http.MethodGet,
		URL:    url,
	}
}
func NewRequestWithBody(
	url string, body io.Reader,
) *xhttp.Request {
	return &xhttp.Request{
		Method: http.MethodGet,
		URL:    url,
		Body:   body,
	}
}
func NewRequestWithMethodAndBody(
	url string, method string,
	body io.Reader,
) *xhttp.Request {
	return &xhttp.Request{
		Method: method,
		URL:    url,
		Body:   body,
	}
}
func NewRequestWithHeaders(
	url string,
	headers map[string][]string,
) *xhttp.Request {
	return &xhttp.Request{
		Method:  http.MethodGet,
		URL:     url,
		Headers: headers,
	}
}
