package http9

import (
	"http_request/pkg/xhttp"
	"net/http"
)

var defaultClient interface {
	Do(r *xhttp.Request) (*http.Response, error)
}

type Request xhttp.Request

func Init(r *Request) error {
	return r.validateAndFillMissing()
}
func (r *Request) validateAndFillMissing() error {
	return nil
}
func (r *Request) Init() error {
	return r.validateAndFillMissing()
}
func Do(r *Request) (*http.Response, error) {
	err := r.Init()
	if err != nil {
		return nil, err
	}

	return defaultClient.Do((*xhttp.Request)(r))
}

func (r *Request) Get() (*http.Response, error) {
	r.Method = http.MethodGet
	err := r.Init()
	if err != nil {
		return nil, err
	}
	return defaultClient.Do((*xhttp.Request)(r))
}
func (r *Request) Post() (*http.Response, error) {
	r.Method = http.MethodGet

	return defaultClient.Do((*xhttp.Request)(r))
}
