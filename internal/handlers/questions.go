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

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var questions []models.Question

	result := database.DB.Find(&questions)
	if result.Error != nil {
		log.Printf("ошибка бд: %v", result.Error)
		http.Error(w, "ошибка при получении данных", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var question models.Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, "неверный json", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(question.Text) == "" {
		http.Error(w, "поле текст не может быть пустым", http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&question)
	if result.Error != nil {
		log.Printf("ошибка бд: %v", result.Error)
		http.Error(w, "ошибка при создании записи", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

func ShowQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	IDStr := vars["id"]

	if _, err := strconv.Atoi(IDStr); err != nil {
		http.Error(w, "введите число", http.StatusBadRequest)
		return
	}

	var question models.Question

	result := database.DB.Preload("Answers").First(&question, IDStr)
	if result.Error != nil {
		log.Printf("ошибка бд: %v", result.Error)
		http.Error(w, "вопрос не найден", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(question)

}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	IDStr := vars["id"]

	result := database.DB.Delete(&models.Question{}, IDStr)
	if result.Error != nil {
		log.Printf("ошибка бд: %v", result.Error)
		http.Error(w, "ошибка при удалении", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "вопрос не найден", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
