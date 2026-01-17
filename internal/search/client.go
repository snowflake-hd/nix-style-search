package search

import (
	"net/http"
	internal "nix-style-search/internal"
)

var HttpClient = &http.Client{}

func ExecuteRequest(req *http.Request) (*http.Response, error) {
	addRequestHeaders(req)

	return HttpClient.Do(req)
}

func addRequestHeaders(req *http.Request) {
	req.SetBasicAuth(internal.SEARCH_USERNAME, internal.SEARCH_PASSWORD)
	req.Header.Set("Content-Type", "application/json")
}
