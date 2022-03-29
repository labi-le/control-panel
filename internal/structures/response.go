package structures

import "time"

// Response is a struct that contains the response from the server.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`

	Data    interface{} `json:"data"`
	Version string      `json:"version"`

	Time time.Time `json:"time"`
}
