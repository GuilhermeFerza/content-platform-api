package main

import (
	"encoding/json"
	"log"
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
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", getTasks)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
