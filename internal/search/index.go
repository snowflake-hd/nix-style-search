package search

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"nix-style-search/internal"
)

type Alias struct {
	Alias         string `json:"alias"`
	Index         string `json:"index"`
	Filter        string `json:"filter"`
	Routing       string `json:"routing"`
	RoutingSearch string `json:"routing.search"`
	IsWriteIndex  string `json:"is_write_index"`
}

var AliasEndpoint = internal.ALIAS_ENDPOINT

func GetAvailableIndices() ([]Alias, error) {
	req, err := http.NewRequest("GET", AliasEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	resp, err := ExecuteRequest(req)
	if err != nil {
		return nil, fmt.Errorf("request indices: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("alias endpoint status %d: %s", resp.StatusCode, string(body))
	}

	var aliases []Alias
	if err := json.Unmarshal(body, &aliases); err == nil {
		return aliases, nil
	}

	aliases, err = parseCatAliases(body)
	if err != nil {
		return nil, fmt.Errorf("parse aliases: %w; body: %s", err, string(body))
	}
	return aliases, nil
}

func parseCatAliases(body []byte) ([]Alias, error) {
	fields := strings.Fields(string(body))
	const cols = 6
	if len(fields)%cols != 0 {
		return nil, fmt.Errorf("unexpected columns: got %d fields", len(fields))
	}
	out := make([]Alias, 0, len(fields)/cols)
	for i := 0; i < len(fields); i += cols {
		out = append(out, Alias{
			Alias:         fields[i],
			Index:         fields[i+1],
			Filter:        fields[i+2],
			Routing:       fields[i+3],
			RoutingSearch: fields[i+4],
			IsWriteIndex:  fields[i+5],
		})
	}
	return out, nil
}
