package api

import (
	"encoding/json"
	"errors"
	"github.com/labi-le/control-panel/internal/structures"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	w := httptest.NewRecorder()
	SuccessResponse(w, "test", "test")

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// extract the response body
	r := structures.Response{}
	if json.Unmarshal(w.Body.Bytes(), &r) != nil {
		t.Errorf("Expected response body to be a valid json, got %s", w.Body.String())
	}

	if r.Success != true {
		t.Errorf("Expected status %s, got %v", "success", r.Success)
	}

	if r.Message != "test" {
		t.Errorf("Expected message %s, got %s", "test", r.Message)
	}

}

func TestBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	BadRequest(w, errors.New("test"))

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	// extract the response body
	r := structures.Response{}
	if json.Unmarshal(w.Body.Bytes(), &r) != nil {
		t.Errorf("Expected response body to be a valid json, got %s", w.Body.String())
	}

	if r.Success != false {
		t.Errorf("Expected status %s, got %v", "success", r.Success)
	}

	if r.Message != "test" {
		t.Errorf("Expected message %s, got %s", "test", r.Message)
	}
}

func TestMethodNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	MethodNotFound(w)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, w.Code)
	}

	if w.Body.String() != "" {
		t.Errorf("Expected empty body, got %s", w.Body.String())
	}
}
