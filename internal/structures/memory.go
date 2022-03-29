package structures

// Memory in kibibyte
type Memory struct {
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}
