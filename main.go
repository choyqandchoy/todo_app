package main

import (
	"log"
	"net/http"
	"todo-api/handlers"
	"todo-api/models"

	"github.com/gorilla/mux"
)

func main() {
	err := models.ConnectDB()
	if err != nil {
		log.Fatal("❌ DB connection error:", err)
	}
	log.Println("✅ Connected to DB")

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("To-Do API is working"))
	})

	r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	secured := r.PathPrefix("/").Subrouter()
	secured.Use(handlers.AuthMiddleware)

	secured.HandleFunc("/tasks", handlers.CreateTaskHandler).Methods("POST")
	secured.HandleFunc("/tasks", handlers.GetTasksHandler).Methods("GET")
	secured.HandleFunc("/tasks/{id}", handlers.GetTaskByIDHandler).Methods("GET")

	log.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
