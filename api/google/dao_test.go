package google

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"searchspring.com/innerspring/common"
	"searchspring.com/innerspring/commontest"
)

func TestCheckUserLoggedIn(t *testing.T) {
	response := &http.Response{
		Body:       ioutil.NopCloser(strings.NewReader(`{"email":"itworked@example.com"}`)),
		StatusCode: 200,
	}
	DAO := NewDAO(common.NewClient(&commontest.MockHttpClient{
		Response: response,
	}))
	email, err := DAO.CheckUserLoggedIn("my token")
	require.Nil(t, err)
	require.Equal(t, email, "itworked@example.com")
}

func TestCheckUserLoggedInErrors(t *testing.T) {

	DAO := NewDAO(common.NewClient(&commontest.MockHttpClient{
		Response: &http.Response{Body: ioutil.NopCloser(&commontest.MockErrorReader{})},
	}))
	_, err := DAO.CheckUserLoggedIn("my token")
	require.NotNil(t, err)
	require.Contains(t, err.Error(), "network error")

	response := &http.Response{
		Body:       ioutil.NopCloser(strings.NewReader(`bad json`)),
		StatusCode: 200,
	}
	DAO = NewDAO(common.NewClient(&commontest.MockHttpClient{
		Response: response,
	}))
	_, err = DAO.CheckUserLoggedIn("my token")
	require.Contains(t, err.Error(), "invalid character")
}
