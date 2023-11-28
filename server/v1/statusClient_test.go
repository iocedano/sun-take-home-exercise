package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestDoRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	mockedClient := StatusClient{}
	mockedClient.client = ts.Client()

	t.Run("should get an StatusResponse", func(t *testing.T) {
		statusCh := make(chan StatusResponse)
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			mockedClient.DoRequest(ts.URL, statusCh, &wg)
		}()

		go func() {
			wg.Wait()
			close(statusCh)
		}()

		result := <-statusCh
		if result.URL != ts.URL {
			t.Fatal("should match the URL")
		}
	})
}

func TestDoRequestBadResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Hello, client")
	}))
	mockedClient := StatusClient{}
	mockedClient.client = ts.Client()

	t.Run("should get an bad request status code", func(t *testing.T) {
		statusCh := make(chan StatusResponse)
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			mockedClient.DoRequest(ts.URL, statusCh, &wg)
		}()

		go func() {
			wg.Wait()
			close(statusCh)
		}()

		result := <-statusCh

		if result.StatusCode != http.StatusBadRequest {
			t.Fatal("should get bad request status code if the response is wrong")
		}
	})
}
