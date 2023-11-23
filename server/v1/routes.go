package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// SendJSONResponse write the commos segment for a JSON response
func sendJSONResponse(w http.ResponseWriter, jsonResp []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func GetSourcesStatusHandler(c StatusClient, sources []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := make([]StatusResponse, 0, len(sources))
		var wg sync.WaitGroup

		wg.Add(len(sources))

		for _, source := range sources {
			go func(url string) {
				request, err := http.NewRequest(http.MethodGet, url, nil)
				if err != nil {
					log.Fatal(err)
				}
				status, err := c.GetStatus(request)
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

func GetSourceStatusHandler(c StatusClient, source string) func(w http.ResponseWriter, r *http.Request) {
	request, err := http.NewRequest(http.MethodGet, source, nil)
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			log.Fatal(err)
		}
		status, _ := c.GetStatus(request)
		jsonResp, err := json.Marshal(status)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		sendJSONResponse(w, jsonResp)
	}
}
