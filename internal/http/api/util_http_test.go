package api

import (
	"encoding/json"
	"github.com/labi-le/control-panel/internal/structures"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	w := httptest.NewRecorder()
	SuccessResponse(w, "test", structures.Response{
		Message: "test",
		Data:    nil,
	})

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// extract the response body
	r := structures.Response{}
	if json.Unmarshal(w.Body.Bytes(), &r) != nil {
		t.Errorf("Expected response body to be a valid json, got %s", w.Body.String())
	}

}
