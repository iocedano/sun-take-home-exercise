package v2

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetHandler returns a HandlerFunc that writes a json status response
func GetHandler(statusClient *StatusClient, sources []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hasMultStatus := len(sources) > 1
		clientResponses := GetStatusFromListSources(sources, statusClient)
		var jsonResp []byte
		var err error

		if hasMultStatus {
			jsonResp, err = json.Marshal(clientResponses)
		} else {
			jsonResp, err = json.Marshal(clientResponses[0])
		}

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}

func SetupRouters(mux *http.ServeMux, sources map[string]string) {
	var listOfUrl []string
	client := &http.Client{}
	statusClient := StatusClient{client: client}

	for name, source := range sources {
		mux.HandleFunc("/v1/"+name+"-status", GetHandler(&statusClient, []string{source}))
		listOfUrl = append(listOfUrl, source)
	}
	mux.HandleFunc("/v1/all-status", GetHandler(&statusClient, listOfUrl))
}
