package structures

// Memory in kibibyte
type Memory struct {
	Total  uint64 `json:"total"`
	Free   uint64 `json:"free"`
	Used   uint64 `json:"used"`
	Cached uint64 `json:"cached"`
}
