package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"searchspring.com/innerspring/model"
)

func TestSignedInUserCheck(t *testing.T) {
	calls := 0
	r := httptest.NewRequest("GET", "http://t", nil)
	wrapSignedInUserCheck(func(authorizationToken string) (string, error) {
		return "example@searchspring.com", nil
	}, func(w http.ResponseWriter, r *http.Request, u *model.User) {
		require.Equal(t, "example@searchspring.com", u.Email)
		calls++
	})(nil, r)
	require.Equal(t, 1, calls)
}
func TestBadUser(t *testing.T) {
	r := httptest.NewRequest("GET", "http://t", nil)
	w := httptest.NewRecorder()
	wrapSignedInUserCheck(func(authorizationToken string) (string, error) {
		return "", fmt.Errorf("bad response")
	}, nil)(w, r)
	require.Equal(t, 403, w.Result().StatusCode)

	w = httptest.NewRecorder()
	os.Setenv("DOMAIN", "searchspring.com")
	wrapSignedInUserCheck(func(authorizationToken string) (string, error) {
		return "notsearchspring@example.com", nil
	}, nil)(w, r)
	require.Equal(t, 403, w.Result().StatusCode)
}

func TestCreateRouter(t *testing.T) {
	_, err := CreateRouter()
	require.Nil(t, err)
}

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "http://localhost:3000/api/projects", ioutil.NopCloser(strings.NewReader(`[{ "name": "Team Lead" }]`)))
	Handler(w, r)
	require.Equal(t, 403, w.Result().StatusCode)
	require.Contains(t, w.Body.String(), "no authorization header")
}
