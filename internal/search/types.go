package search

type SearchRequest struct {
	From  int                    `json:"from"`
	Size  int                    `json:"size"`
	Sort  []map[string]string    `json:"sort"`
	Aggs  map[string]interface{} `json:"aggs"`
	Query map[string]interface{} `json:"query"`
}
