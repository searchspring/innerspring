package commontest

import (
	"fmt"
	"net/http"
)

type MockHttpClient struct {
	Error    error
	Response *http.Response
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.Response, m.Error
}

func newMockHttpClient(response *http.Response, err error) *MockHttpClient {
	return &MockHttpClient{
		Response: response,
		Error:    err,
	}
}

type MockErrorReader struct {
}

func (*MockErrorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("network error")
}
