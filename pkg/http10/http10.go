package http10

import (
	"errors"
	"http_request/pkg/xhttp"
	"net/http"
)

var defaultClient interface {
	Do(r *xhttp.Request) (*http.Response, error)
}

type Request xhttp.Request

func (r *Request) init() error {
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	if r.URL == "" {
		return errors.New("url required")
	}
	return nil
}
func (r Request) Get(url string) (*http.Response, error) {
	r.URL = url
	r.Method = http.MethodGet

	if err := r.init(); err != nil {
		return nil, err
	}

	if err := r.init(); err != nil {
		return nil, err
	}

	return defaultClient.Do((*xhttp.Request)(&r))
}
func (r Request) Post(url string) (*http.Response, error) {
	r.URL = url
	r.Method = http.MethodPost

	if err := r.init(); err != nil {
		return nil, err
	}

	return defaultClient.Do((*xhttp.Request)(&r))
}
