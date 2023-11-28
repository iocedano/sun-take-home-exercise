package v1

import (
	"net/http"
	"sync"
	"testing"
)

type MockedStatusClient struct {
	StatusClient
}

func (MockedStatusClient) DoRequest(source string, statusCha chan StatusResponse, wg *sync.WaitGroup) {
	go func() {
		statusCha <- StatusResponse{URL: source}
		wg.Done()
	}()
}

func (MockedStatusClient) Get(request *http.Request) (StatusResponse, error) {
	return StatusResponse{
		URL: request.URL.String(),
	}, nil
}

func TestGetStatus(t *testing.T) {
	mockedClient := MockedStatusClient{}

	t.Run("should match the source", func(t *testing.T) {
		expected := StatusResponse{URL: "test"}
		result := GetStatusFromListSources([]string{"test"}, &mockedClient)

		if result[0] != expected {
			t.Fatalf("It doesn't match the expected result; expect %v and got %v", expected, result)
		}
	})

	t.Run("shoudl got empty list from an empty array", func(t *testing.T) {
		mockedClient := MockedStatusClient{}
		result := GetStatusFromListSources([]string{}, &mockedClient)

		if len(result) != 0 {
			t.Fatalf("It expected an empty list result; but got %v", result)
		}
	})
}
