package http7

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestOptions(t *testing.T) {
	require.Equal(t,
		[]option{
			WithMethod(http.MethodGet),
			WithHeader("foo", "bar", "baz"),
		},
		[]option{
			WithMethod(http.MethodGet),
			WithHeader("foo", "bar", "baz"),
		},
	)
}
