package common

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteError(t *testing.T) {
	w := httptest.NewRecorder()
	WriteError(w, 500, "some kinda error")
	require.Equal(t, 500, w.Result().StatusCode)
	require.Contains(t, w.Body.String(), "some kinda error")
}

type recorder struct{}

func (*recorder) Write([]byte) (int, error)  { return -1, fmt.Errorf("badness") }
func (*recorder) WriteHeader(statusCode int) {}
func (*recorder) Header() http.Header        { return nil }

func TestError(t *testing.T) {
	WriteError(&recorder{}, 500, "some kinda error")

}
