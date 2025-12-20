package main

import (
	"log"
	"net/http"

	"github.com/GuilhermeFerza/content-platform-api/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", handler.Tasks)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
