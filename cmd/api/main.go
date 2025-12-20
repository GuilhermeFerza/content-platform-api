package main

import (
	"log"
	"net/http"

	"github.com/GuilhermeFerza/content-platform-api/internal/handler"
	"github.com/GuilhermeFerza/content-platform-api/internal/service"
)

func main() {
	mux := http.NewServeMux()

	taskService := service.NewTaskService()
	taskHandler := handler.NewTaskHandler(taskService)

	mux.HandleFunc("/tasks", taskHandler.Tasks)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
