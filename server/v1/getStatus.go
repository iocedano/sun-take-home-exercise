package v1

import (
	"net/http"
	"time"
)

type StatusResponse struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duraction  int    `json:"duration"`
	Date       int64  `json:"date"`
}

type StatusClient interface {
	GetStatus(request *http.Request) (StatusResponse, error)
}

type statusHTTPClient struct {
	client *http.Client
}

// GetStatus do a client request and return the status response struct
// StatusHTTPClient has the GetStatus method which implements the StatusClient interface
func (s statusHTTPClient) GetStatus(request *http.Request) (StatusResponse, error) {
	startTime := time.Now()
	resp, err := s.client.Do(request)
	duration := time.Since(startTime)

	if err != nil {
		return StatusResponse{Url: request.URL.String(), StatusCode: http.StatusBadRequest}, err
	}

	defer resp.Body.Close()

	return StatusResponse{
		Url:        request.URL.String(),
		StatusCode: resp.StatusCode,
		Duraction:  int(duration.Milliseconds()),
		Date:       time.Now().Unix(),
	}, nil
}

// NewStatusClient return a status HTTP Client instance
// Keeping the statusHTTPClient privite in the package
func NewStatusClient(c *http.Client) *statusHTTPClient {
	return &statusHTTPClient{
		client: c,
	}
}
