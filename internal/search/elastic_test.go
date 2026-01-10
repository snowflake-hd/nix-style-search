package search

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestBuildSearchQuery(t *testing.T) {
	req, err := BuildSearchQuery("test", 0, 10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	size, ok := req["size"].(int)
	if !ok || size != 10 {
		t.Fatalf("expected size 10, got %v", req["size"])
	}
	from, ok := req["from"].(int)
	if !ok || from != 0 {
		t.Fatalf("expected from 0, got %v", req["from"])
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(jsonData), `"query":"test"`) {
		t.Errorf("expected query string 'test' to be in the JSON output")
	}
}
