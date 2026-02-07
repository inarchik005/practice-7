package main

import (
	"log"
	"net/http"
	"task-api/internal/handlers"
	"task-api/internal/middleware"
	"task-api/internal/storage"
)

func main() {
	store := storage.NewTaskStore()
	handler := handlers.NewTaskHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", handler.GetTasks)
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("PATCH /tasks", handler.UpdateTask)

	stack := middleware.APIKeyMiddleware(
		middleware.LoggingMiddleware(mux),
	)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", stack))
}
