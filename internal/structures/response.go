package structures

// Response is a struct that contains the response from the server.
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
