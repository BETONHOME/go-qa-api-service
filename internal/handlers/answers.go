package handlers

import (
	"API_T3/internal/database"
	"API_T3/internal/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	questionIDStr := vars["id"]

	questionID, err := strconv.ParseUint(questionIDStr, 10, 32)
	if err != nil {
		http.Error(w, "неверный ID вопроса", http.StatusBadRequest)
		return
	}

	var answer models.Answer
	err = json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(w, "неверный json", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(answer.Text) == "" {
		http.Error(w, "ответ не может быть пустым", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(answer.UserID) == "" {
		http.Error(w, "введите имя", http.StatusBadRequest)
		return
	}

	answer.QuestionID = uint(questionID)

	result := database.DB.Create(&answer)
	if result.Error != nil {
		log.Printf("ошибка бд: %v", result.Error)
		http.Error(w, "ошибка при создании ответа", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

func ShowAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	IDStr := vars["id"]

	if _, err := strconv.Atoi(IDStr); err != nil {
		http.Error(w, "введите число", http.StatusBadRequest)
		return
	}

	var answer models.Answer
	result := database.DB.First(&answer, IDStr)
	if result.Error != nil {
		log.Printf("ошибка бд %s: %v", IDStr, result.Error)
		http.Error(w, "ответ не найден", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(answer)
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	IDStr := vars["id"]

	result := database.DB.Delete(&models.Answer{}, IDStr)
	if result.Error != nil {
		log.Printf("ошибка бд %s: %v", IDStr, result.Error)
		http.Error(w, "ошибка при удалении", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
