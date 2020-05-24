package common

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"searchspring.com/innerspring/commontest"
)

func TestHappyPath(t *testing.T) {
	response := &http.Response{
		Body:       ioutil.NopCloser(strings.NewReader("everything worked out")),
		StatusCode: 200,
	}
	body, err := NewClient(&commontest.MockHttpClient{
		Response: response,
	}).AuthorizedGet("token", "http://example.com")
	require.Nil(t, err)
	require.Equal(t, "everything worked out", string(body))
}

func TestAuthorizedRequestErrors(t *testing.T) {
	_, err := NewClient(nil).AuthorizedGet("", "")
	require.Contains(t, err.Error(), "no auth")

	_, err = NewClient(nil).AuthorizedGet("token", ":::")
	require.Contains(t, err.Error(), "missing protocol scheme")

	_, err = NewClient(&commontest.MockHttpClient{
		Error: fmt.Errorf("not implemented"),
	}).AuthorizedGet("token", "http://example.com")
	require.Contains(t, err.Error(), "not implemented")

	_, err = NewClient(&commontest.MockHttpClient{
		Response: &http.Response{Body: ioutil.NopCloser(&commontest.MockErrorReader{})},
	}).AuthorizedGet("token", "http://example.com")
	require.Contains(t, err.Error(), "network error")

	response := &http.Response{
		Body:       ioutil.NopCloser(strings.NewReader("google is down")),
		StatusCode: 500,
	}
	_, err = NewClient(&commontest.MockHttpClient{
		Response: response,
	}).AuthorizedGet("token", "http://example.com")
	require.Contains(t, err.Error(), "google is down")
}
