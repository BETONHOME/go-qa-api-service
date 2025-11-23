package main

import (
	"API_T3/internal/database"
	"API_T3/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.InitDB()

	router := mux.NewRouter()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Started %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/questions", http.StatusSeeOther)
	})

	router.HandleFunc("/questions", handlers.GetQuestions).Methods("GET")
	router.HandleFunc("/questions", handlers.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", handlers.ShowQuestion).Methods("GET")
	router.HandleFunc("/questions/{id}", handlers.DeleteQuestion).Methods("DELETE")

	router.HandleFunc("/questions/{id}/answers", handlers.CreateAnswer).Methods("POST")
	router.HandleFunc("/answers/{id}", handlers.ShowAnswer).Methods("GET")
	router.HandleFunc("/answers/{id}", handlers.DeleteAnswer).Methods("DELETE")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal("error:", err)
	}
}
