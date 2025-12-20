package handler

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Aprender Go", Done: false},
	{ID: 2, Title: "Criar uma API", Done: true},
	{ID: 3, Title: "Teste", Done: true},
}

func Tasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var newTask Task

		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if newTask.Title == "" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}

		newTask.ID = len(tasks) + 1

		tasks = append(tasks, newTask)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
