package v2

import (
	"sync"
)

// GetStatusFromListSources return a list of StatusResponse struct from a list of souces
// using the statusClients methods
func GetStatusFromListSources(listOfSoucers []string, statusClient *StatusClient) []StatusResponse {
	statusRespCh := make(chan StatusResponse)
	clientResponses := make([]StatusResponse, 0, len(listOfSoucers))
	var wg sync.WaitGroup

	wg.Add(len(listOfSoucers))

	for _, source := range listOfSoucers {
		statusClient.DoRequest(source, statusRespCh, &wg)
	}

	go func() {
		wg.Wait()
		close(statusRespCh)
	}()

	for resp := range statusRespCh {
		clientResponses = append(clientResponses, resp)
	}

	return clientResponses
}
