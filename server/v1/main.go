package v1

import "net/http"

func SetupRouters(mux *http.ServeMux, sources map[string]string) {
	client := &http.Client{}
	var listOfUrl []string

	sc := NewStatusClient(client)
	// Rest API
	for name, source := range sources {
		mux.HandleFunc("/v1/"+name+"-status", GetSourceStatusHandler(sc, source))
		listOfUrl = append(listOfUrl, source)
	}
	mux.HandleFunc("/v1/all-status", GetSourcesStatusHandler(sc, listOfUrl))
}
