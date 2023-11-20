package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"sync"
	"time"
)

var sources = map[string]string{
	"google": "https://www.google.com",
	"amazon": "https://www.amazon.com",
}

type StatusResponse struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duraction  int    `json:"duration"`
	Date       int64  `json:"date"`
}

func main() {
	mux := http.NewServeMux()
	var listOfUrl []string

	// Client
	buildPath := path.Clean("client/build")
	mux.Handle("/", http.FileServer(http.Dir(buildPath)))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("client/build/static"))))

	// Rest API
	for name, source := range sources {
		mux.HandleFunc("/v1/"+name+"-status", GetHandlerSourceStatus(source))
		listOfUrl = append(listOfUrl, source)
	}

	mux.HandleFunc("/v1/all-status", GetHandlerSourcesStatus(listOfUrl))

	http.ListenAndServe(":8080", mux)
}

func GetStatus(source string) (StatusResponse, error) {
	startTime := time.Now()
	resp, err := http.Get(source)
	duration := time.Since(startTime)

	if err != nil {
		return StatusResponse{Url: source, StatusCode: http.StatusBadRequest}, err
	}

	defer resp.Body.Close()

	return StatusResponse{
		Url:        source,
		StatusCode: resp.StatusCode,
		Duraction:  int(duration.Milliseconds()),
		Date:       time.Now().Unix(),
	}, nil
}

func sendJSONResponse(w http.ResponseWriter, jsonResp []byte) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func GetHandlerSourcesStatus(sources []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := make([]StatusResponse, 0, len(sources))
		var wg sync.WaitGroup

		wg.Add(len(sources))

		for _, source := range sources {
			go func(url string) {
				status, err := GetStatus(url)
				if err != nil {
					log.Fatalf("Error happened in Getting the status. Err: %s", err)
				}
				resp = append(resp, status)
				defer wg.Done()
			}(source)
		}

		wg.Wait()

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		sendJSONResponse(w, jsonResp)
	}
}

func GetHandlerSourceStatus(source string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status, _ := GetStatus(source)
		jsonResp, err := json.Marshal(status)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		sendJSONResponse(w, jsonResp)
	}
}
