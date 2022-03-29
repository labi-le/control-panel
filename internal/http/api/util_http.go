package api

import (
	"encoding/json"
	"github.com/labi-le/control-panel/internal/structures"
	"net/http"
	"time"
)

func SuccessResponse(w http.ResponseWriter, msg string, data interface{}) {
	r := structures.Response{
		Success: true,
		Message: msg,
		Data:    data,
	}

	response(w, r)
}

func BadRequest(w http.ResponseWriter, err error) {
	r := structures.Response{
		Success: false,
		Message: err.Error(),
		Data:    []string{},
	}

	response(w, r)
}

func MethodNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func response(w http.ResponseWriter, response structures.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	response.Time = time.Now()
	response.Version = "1.0.0"

	switch response.Success {
	case false:
		w.WriteHeader(http.StatusBadRequest)

	case true:
		w.WriteHeader(http.StatusOK)
	}

	_ = json.NewEncoder(w).Encode(response)
}
