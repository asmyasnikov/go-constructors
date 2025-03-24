package main

import (
	"io"
	"net/http"
	"strings"

	"http_request/pkg/http1"
	"http_request/pkg/http10"
	"http_request/pkg/http2"
	"http_request/pkg/http3"
	"http_request/pkg/http4"
	"http_request/pkg/http5"
	"http_request/pkg/http6"
	"http_request/pkg/http7"
	"http_request/pkg/http8"
	"http_request/pkg/http9"
	"http_request/pkg/xhttp"
)

func explicitConstructor() {
	_ = &xhttp.Request{
		Method: http.MethodGet,
		URL:    "http://example.com",
		Body: io.NopCloser(
			strings.NewReader("some body"),
		),
		Headers: map[string][]string{
			"foo": []string{"bar", "baz"},
		},
	}
}

func http1Constructor() {
	_ = http1.NewRequestWithMethodAndBody(
		"http://example.com",
		http.MethodGet,
		strings.NewReader("some body"),
	)
}

func http2Constructor() {
	_ = http2.NewRequest3(
		"http://example.com",
		http.MethodGet,
		strings.NewReader("some body"),
	)
}

func http3ConstructorWithElipsis() {
	_, _ = http3.NewRequest(
		"http://example.com",
		http.MethodGet,
		strings.NewReader("some body"),
	)
}

func http4ConstructorWithSettings() {
	_ = http4.NewRequest(
		"http://example.com",
		&http4.Settings{
			Method: http.MethodGet,
			Body:   "",
		},
	)
}

func http5ConstructorWithSettings() {
	body := "{}"
	_ = http5.NewRequest(
		"http://example.com",
		&http5.Settings{
			Body: &body,
		},
	)
}

func ptr[T any](v T) *T {
	return &v
}

func http5ConstructorWithSettingsPtr() {
	_ = http5.NewRequest(
		"http://example.com",
		&http5.Settings{
			Body: ptr("{}"),
		},
	)
}

func http6ConstructorWithFuncOpts() {
	_, _ = http6.NewRequest(
		"http://example.com",
		http6.WithMethod(http.MethodPost),
		http6.WithHeader("foo", "bar", "baz"),
	)
}

func http7ConstructorWithOpts() {
	_, _ = http7.NewRequest(
		"http://example.com",
		http7.WithMethod(http.MethodPost),
		http7.WithBody(strings.NewReader("some body")),
	)
}

func http8ConstructorWithBuilder() {
	_, _ = http8.NewRequest("http://example.com").
		WithMethod(http.MethodPost).
		WithBody(strings.NewReader("some body")).
		Build()
}
func http9ConstructorWithPublicInit() {
	r := http9.Request{
		URL: "http://example.com",
	}
	if err := r.Init(); err != nil {
		panic(err)
	}
}
func http9ConstructorWithPublicInit2() {
	r := &http9.Request{
		URL: "http://example.com",
	}
	if err := http9.Init(r); err != nil {
		panic(err)
	}
}
func http9ConstructorWithPublicInit3() {
	r := &http9.Request{
		URL: "http://example.com",
	}
	_, err := r.Get()
	if err != nil {
		panic(err)
	}
}
func http9ConstructorWithPublicInit4() {
	r := &http9.Request{
		URL: "http://example.com",
	}
	_, err := http9.Do(r)
	if err != nil {
		panic(err)
	}
}
func http10ConstructorWithPrivateInit() {
	var r http10.Request
	_, _ = r.Get("http://example.com")
}

func main() {
	_ = http1.NewRequest("http://example.com")
	_ = http1.NewRequestWithMethodAndBody("http://example.com",
		http.MethodPost,
		strings.NewReader("some body"),
	)
	_ = http2.NewRequest("http://example.com")
	_ = http2.NewRequest3("http://example.com",
		http.MethodPost,
		strings.NewReader("some body"),
	)
	_, _ = http3.NewRequest("http://example.com")
	_, _ = http3.NewRequest("http://example.com",
		http.MethodPost,
		strings.NewReader("some body"),
	)
	_ = http4.NewRequest("http://example.com", nil)
	_ = http4.NewRequest("http://example.com", &http4.Settings{
		Method: http.MethodPost,
		Body:   "some body",
	})
	_ = http5.NewRequest("http://example.com", nil)
	_ = http5.NewRequest("http://example.com", &http5.Settings{
		Method: ptr(http.MethodPost),
		Body:   ptr("some body"),
	})
	_, _ = http6.NewRequest("http://example.com")
	_, _ = http6.NewRequest("http://example.com",
		http6.WithMethod(http.MethodPost),
		http6.WithBody(strings.NewReader("some body")),
	)
	_, _ = http7.NewRequest("http://example.com")
	_, _ = http7.NewRequest("http://example.com",
		http7.WithMethod(http.MethodPost),
		http7.WithBody(strings.NewReader("some body")),
	)
	_, _ = http8.NewRequest("http://example.com").
		Build()
	_, _ = http8.NewRequest("http://example.com").
		WithMethod(http.MethodPost).
		WithBody(strings.NewReader("some body")).
		Build()
	r9 := http9.Request{
		URL: "http://example.com",
	}
	if err := r9.Init(); err != nil {
		panic(err)
	}
	_, _ = r9.Get()
	var r10 http10.Request
	_, _ = r10.Get("http://example.com")
	r10_1 := http10.Request{
		Headers: map[string][]string{
			"foo": {"bar", "baz"},
		},
	}
	_, _ = r10_1.Get("http://example.com")
}
