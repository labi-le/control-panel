package api

import (
	"encoding/json"
	"github.com/labi-le/control-panel/internal"
	"github.com/labi-le/control-panel/internal/structures"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func SuccessResponse(w *echo.Response, d any) {
	response(w, d)

	w.WriteHeader(http.StatusOK)
}

func BadRequest(w *echo.Response, err error) {
	r := structures.Response{
		Message: err.Error(),
		Data:    []string{},
	}

	response(w, r)

	w.WriteHeader(http.StatusBadRequest)
}

func response(w *echo.Response, r any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Version", internal.PanelVersion)
	w.Header().Set("Date", time.Now().String())

	if err := json.NewEncoder(w).Encode(r); err != nil {
		panic(err)
	}
}
