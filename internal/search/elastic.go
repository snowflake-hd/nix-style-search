package search

import (
	"encoding/json"
)

// tbh I'm to lazy to make all those nested structs ...

func BuildSearchQuery(queryStr string, from int, size int) (map[string]interface{}, error) {
	query := map[string]interface{}{
		"from": from,
		"size": size,
		"sort": []map[string]string{
			{
				"_score":            "desc",
				"package_attr_name": "desc",
				"package_pversion":  "desc",
			},
		},
		"aggs": map[string]interface{}{
			"package_attr_set": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "package_attr_set",
					"size":  20,
				},
			},
			"package_license_set": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "package_license_set",
					"size":  20,
				},
			},
			"package_maintainers_set": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "package_maintainers_set",
					"size":  20,
				},
			},
			"package_teams_set": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "package_teams_set",
					"size":  20,
				},
			},
			"package_platforms": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "package_platforms",
					"size":  20,
				},
			},
			"all": map[string]interface{}{
				"global": map[string]interface{}{},
				"aggs": map[string]interface{}{
					"package_attr_set": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "package_attr_set",
							"size":  20,
						},
					},
					"package_license_set": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "package_license_set",
							"size":  20,
						},
					},
					"package_maintainers_set": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "package_maintainers_set",
							"size":  20,
						},
					},
					"package_teams_set": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "package_teams_set",
							"size":  20,
						},
					},
					"package_platforms": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "package_platforms",
							"size":  20,
						},
					},
				},
			},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"type": map[string]interface{}{
								"value": "package",
								"_name": "filter_packages",
							},
						},
					},
					map[string]interface{}{
						"bool": map[string]interface{}{
							"must": []interface{}{
								map[string]interface{}{"bool": map[string]interface{}{"should": []interface{}{}}},
								map[string]interface{}{"bool": map[string]interface{}{"should": []interface{}{}}},
								map[string]interface{}{"bool": map[string]interface{}{"should": []interface{}{}}},
								map[string]interface{}{"bool": map[string]interface{}{"should": []interface{}{}}},
								map[string]interface{}{"bool": map[string]interface{}{"should": []interface{}{}}},
							},
						},
					},
				},
				"must": []interface{}{
					map[string]interface{}{
						"dis_max": map[string]interface{}{
							"tie_breaker": 0.7,
							"queries": []interface{}{
								map[string]interface{}{
									"multi_match": map[string]interface{}{
										"type":                                "cross_fields",
										"query":                               queryStr,
										"analyzer":                            "whitespace",
										"auto_generate_synonyms_phrase_query": false,
										"operator":                            "and",
										"_name":                               "multi_match_test",
										"fields": []string{
											"package_attr_name^9",
											"package_attr_name.*^5.3999999999999995",
											"package_programs^9",
											"package_programs.*^5.3999999999999995",
											"package_pname^6",
											"package_pname.*^3.5999999999999996",
											"package_description^1.3",
											"package_description.*^0.78",
											"package_longDescription^1",
											"package_longDescription.*^0.6",
											"flake_name^0.5",
											"flake_name.*^0.3",
										},
									},
								},
								map[string]interface{}{
									"wildcard": map[string]interface{}{
										"package_attr_name": map[string]interface{}{
											"value":            "*" + queryStr + "*",
											"case_insensitive": true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return query, nil
}

func QueryToJSON(query map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(query, "", "  ")
}
