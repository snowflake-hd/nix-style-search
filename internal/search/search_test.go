package search

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAvailableIndicesSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"alias": "test", "index": "test-index"}]`))
	}))
	defer server.Close()

	originalEndpoint := AliasEndpoint
	AliasEndpoint = server.URL
	defer func() { AliasEndpoint = originalEndpoint }()

	aliases, err := GetAvailableIndices()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(aliases) != 1 || aliases[0].Alias != "test" {
		t.Fatalf("expected one alias, got %v", aliases)
	}
}

func TestGetAvailableIndicesPlainText(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("alias1 index1 - - - -"))
	}))
	defer server.Close()

	originalEndpoint := AliasEndpoint
	AliasEndpoint = server.URL
	defer func() { AliasEndpoint = originalEndpoint }()

	aliases, err := GetAvailableIndices()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(aliases) != 1 || aliases[0].Alias != "alias1" {
		t.Fatalf("expected one alias from plain text, got %v", aliases)
	}
}

func TestReadResponseBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"hits":{"total":{"value":1},"hits":[{"_source":{"package_attr_name":"test"}}]}}`))
	}))
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	searchResp, err := readResponseBody(resp)
	if err != nil {
		t.Fatalf("expected no error from readResponseBody, got %v", err)
	}
	if searchResp.Hits.Total.Value != 1 {
		t.Fatalf("expected 1 hit, got %d", searchResp.Hits.Total.Value)
	}
}
