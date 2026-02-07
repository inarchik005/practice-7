package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task-api/internal/models"
	"task-api/internal/storage"
)

type TaskHandler struct {
	store *storage.TaskStore
}

func NewTaskHandler(store *storage.TaskStore) *TaskHandler {
	return &TaskHandler{store: store}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")

	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
			return
		}

		task, ok := h.store.GetByID(id)
		if !ok {
			http.Error(w, `{"error":"task not found"}`, http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(task)
		return
	}

	json.NewEncoder(w).Encode(h.store.GetAll())
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		http.Error(w, `{"error":"invalid title"}`, http.StatusBadRequest)
		return
	}

	task := h.store.Create(models.Task{
		Title: req.Title,
		Done:  false,
	})

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"invalid id"}`, http.StatusBadRequest)
		return
	}

	var req struct {
		Done bool `json:"done"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"invalid body"}`, http.StatusBadRequest)
		return
	}

	task, ok := h.store.Update(id, req.Done)
	if !ok {
		http.Error(w, `{"error":"task not found"}`, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(task)
}
