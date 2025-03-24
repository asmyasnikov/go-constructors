package xhttp

import (
	"io"
)

type Request struct {
	Method  string
	URL     string
	Body    io.Reader
	Headers map[string][]string
}
