package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateQuestion_EmptyText(t *testing.T) {
	body := `{"text": ""}`
	req := httptest.NewRequest("POST", "/questions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CreateQuestion(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateQuestion_InvalidJSON(t *testing.T) {
	body := `{"text":}`
	req := httptest.NewRequest("POST", "/questions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	CreateQuestion(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestShowQuestion_InvalidID(t *testing.T) {
	req := httptest.NewRequest("GET", "/questions/abc", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/questions/{id}", ShowQuestion)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
