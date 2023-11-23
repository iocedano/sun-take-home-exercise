package v1

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockedStatusHTTPClient struct{}

func (MockedStatusHTTPClient) GetStatus(request *http.Request) (StatusResponse, error) {
	return StatusResponse{Url: request.URL.String()}, nil
}

func TestGetSourcesStatusHandler(t *testing.T) {
	c := MockedStatusHTTPClient{}
	sources := []string{"https://www.test.com", "https://www.test.com"}
	expected := []StatusResponse{{Url: "https://www.test.com"}, {Url: "https://www.test.com"}}
	handler := GetSourcesStatusHandler(c, sources)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	handler(w, r)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	var clientResp []StatusResponse

	if err := json.Unmarshal(body, &clientResp); err != nil {
		t.Error("unexpect unmarshal client json response")
	}

	for i, cliExp := range clientResp {
		if cliExp != expected[i] {
			t.Errorf("unexpected client json response, expected: %v and got %v", expected[i], cliExp)
		}
	}
}

func TestGetSourceStatusHandler(t *testing.T) {
	c := MockedStatusHTTPClient{}
	handler := GetSourceStatusHandler(c, "https://www.test.com")
	expected := StatusResponse{Url: "https://www.test.com"}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	handler(w, r)
	resp := w.Result()

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected content type to be %s got %s", "application/json", contentType)
	}

	body, _ := io.ReadAll(resp.Body)

	var clientResp StatusResponse

	if err := json.Unmarshal(body, &clientResp); err != nil {
		t.Error("unexpect unmarshal client json response")
	}

	if expected != clientResp {
		t.Errorf("unexpected client json response, expected: %v and got %v", expected, clientResp)
	}
}
