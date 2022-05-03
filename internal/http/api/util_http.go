package api

import (
	"encoding/json"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"net/http"
	"time"
)

func SuccessResponse(w http.ResponseWriter, d any) {
	w.WriteHeader(http.StatusOK)
	response(w, d)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)

	r := structures.Response{
		Message: err.Error(),
		Data:    []string{},
	}

	response(w, r)
}

func response(w http.ResponseWriter, r any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Version", internal.PanelVersion)
	w.Header().Set("Date", time.Now().String())

	if err := json.NewEncoder(w).Encode(r); err != nil {
		panic(err)
	}
}
