// Testing
// Use the httptest
// Ensure the response match
// Ensure the the collect the response list from a slice of sources
package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockServer(code int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("TEST")
		if http.StatusOK != code {
			w.WriteHeader(code)
			return
		}
		w.Write([]byte("test"))
	}))
}

func getMockedStatusHTTPClient(svr *httptest.Server) *statusHTTPClient {
	return NewStatusClient(svr.Client())
}

func TestGetStatusSame(t *testing.T) {
	svr := MockServer(http.StatusOK)
	mockedClient := getMockedStatusHTTPClient(svr)
	req := httptest.NewRequest("GET", svr.URL, nil)
	expected := StatusResponse{Url: svr.URL, StatusCode: http.StatusOK, Duraction: 0, Date: 0}
	fmt.Println("TEST", svr.URL)
	defer svr.Close()

	resp, _ := mockedClient.GetStatus(req)

	if resp.Url != expected.Url {
		t.Errorf("expected res to be %s got %s", expected.Url, resp.Url)
	}
}
