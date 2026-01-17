package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"nix-style-search/internal"
)

type Response struct {
	Took     int    `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hits   `json:"hits"`
}

type Shards struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type Hits struct {
	Total    Total    `json:"total"`
	MaxScore *float64 `json:"max_score"`
	Hits     []Hit    `json:"hits"`
}

type Total struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}

type Hit struct {
	Index  string        `json:"_index"`
	Type   string        `json:"_type"`
	ID     string        `json:"_id"`
	Score  *float64      `json:"_score"`
	Source PackageSource `json:"_source"`
	Sort   []interface{} `json:"sort"`
}

type PackageSource struct {
	Type                   string          `json:"type"`
	PackageAttrName        string          `json:"package_attr_name"`
	PackageAttrSet         string          `json:"package_attr_set"`
	PackagePname           string          `json:"package_pname"`
	PackagePversion        string          `json:"package_pversion"`
	PackagePlatforms       []string        `json:"package_platforms"`
	PackageOutputs         []string        `json:"package_outputs"`
	PackageDefaultOutput   string          `json:"package_default_output"`
	PackagePrograms        []string        `json:"package_programs"`
	PackageMainProgram     string          `json:"package_mainProgram"`
	PackageLicense         []License       `json:"package_license"`
	PackageLicenseSet      []string        `json:"package_license_set"`
	PackageMaintainers     []Maintainer    `json:"package_maintainers"`
	PackageMaintainersSet  []string        `json:"package_maintainers_set"`
	PackageTeams           FlexibleStrings `json:"package_teams"`
	PackageTeamsSet        []string        `json:"package_teams_set"`
	PackageDescription     string          `json:"package_description"`
	PackageLongDescription string          `json:"package_longDescription"`
	PackageHydra           interface{}     `json:"package_hydra"`
	PackageSystem          string          `json:"package_system"`
	PackageHomepage        []string        `json:"package_homepage"`
	PackagePosition        string          `json:"package_position"`
}

type License struct {
	URL      string `json:"url"`
	FullName string `json:"fullName"`
}

type Maintainer struct {
	Name   string `json:"name"`
	Github string `json:"github"`
	Email  string `json:"email"`
}

type FlexibleStrings []string

func (fs *FlexibleStrings) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*fs = nil
		return nil
	}

	switch data[0] {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		*fs = []string{s}
		return nil
	case '[':
		var arr []string
		if err := json.Unmarshal(data, &arr); err == nil {
			*fs = arr
			return nil
		}
		var raw []interface{}
		if err := json.Unmarshal(data, &raw); err == nil {
			res := make([]string, 0, len(raw))
			for _, v := range raw {
				res = append(res, fmt.Sprint(v))
			}
			*fs = res
			return nil
		}
		return fmt.Errorf("unexpected array content for FlexibleStrings: %s", string(data))
	case '{':
		var m map[string]interface{}
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}
		res := make([]string, 0, len(m))
		for k, v := range m {
			if s, ok := v.(string); ok && s != "" {
				res = append(res, fmt.Sprintf("%s=%s", k, s))
			} else {
				res = append(res, k)
			}
		}
		*fs = res
		return nil
	default:
		return fmt.Errorf("unexpected JSON type for FlexibleStrings: %s", string(data))
	}
}

func Query(query string, index string, from int, size int) (*Response, error) {
	req, err := createRequest(index, query, from, size)
	if err != nil {
		return nil, fmt.Errorf("failed to create search request: %w", err)
	}

	resp, err := ExecuteRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search request: %w", err)
	}

	return readResponseBody(resp)
}

func readResponseBody(resp *http.Response) (*Response, error) {
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("elasticsearch returned status %d: %s", resp.StatusCode, string(body))
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response JSON: %w", err)
	}

	return &result, nil
}

func createRequest(index string, query string, from int, size int) (*http.Request, error) {
	elasticQuery, err := BuildSearchQuery(query, from, size)

	jsonData, err := QueryToJSON(elasticQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal search request: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(internal.SEARCH_ENDPOINT, index), bytes.NewReader(jsonData))

	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	return req, nil
}
