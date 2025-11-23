package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateAnswer_EmptyText(t *testing.T) {
	body := `{"user_id": "user123", "text": ""}`
	req := httptest.NewRequest("POST", "/questions/1/answers", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/questions/{id}/answers", CreateAnswer)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCreateAnswer_EmptyUserID(t *testing.T) {
	body := `{"user_id": "", "text": "Ответ"}`
	req := httptest.NewRequest("POST", "/questions/1/answers", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/questions/{id}/answers", CreateAnswer)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
