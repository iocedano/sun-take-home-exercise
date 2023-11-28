package v1

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type StatusResponse struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duraction  int    `json:"duraction"`
	Date       int64  `json:"date"`
}

type StatusClient struct {
	client *http.Client
}

type StatusIClient interface {
	Get(request *http.Request) (StatusResponse, error)
	DoRequest(source string, statusCha chan StatusResponse, wg *sync.WaitGroup)
}

func (sc *StatusClient) Get(request *http.Request) (StatusResponse, error) {
	startTime := time.Now()
	resp, err := sc.client.Do(request)
	duration := time.Since(startTime)

	if err != nil {
		return StatusResponse{URL: request.URL.String(), StatusCode: http.StatusBadRequest}, err
	}

	defer resp.Body.Close()

	return StatusResponse{
		URL:        request.URL.String(),
		StatusCode: resp.StatusCode,
		Duraction:  int(duration.Milliseconds()),
		Date:       time.Now().Unix(),
	}, nil
}

func (sc *StatusClient) DoRequest(source string, statusCha chan StatusResponse, wg *sync.WaitGroup) {
	go func(uri string) {
		if source == "" {
			statusCha <- StatusResponse{}
			return
		}
		defer wg.Done()
		request, err := http.NewRequest(http.MethodGet, uri, nil)
		if err != nil {
			log.Fatal(err)
		}
		status, err := sc.Get(request)
		if err != nil {
			log.Fatalf("Error happened in Getting the status. Err: %s", err)
		}
		statusCha <- status
	}(source)
}
