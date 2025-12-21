package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/GuilhermeFerza/content-platform-api/internal/model"
	"github.com/GuilhermeFerza/content-platform-api/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case http.MethodGet:
		tasks := h.service.GetAll()
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		defer r.Body.Close()

		var task model.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if task.Title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}

		created := h.service.Create(task)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h *TaskHandler) TaskByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])

	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodDelete:
		ok := h.service.DeleteByID(id)
		if !ok {
			http.Error(w, "Task not Found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}
