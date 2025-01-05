package server

import (
	"go-web/internal/handlers"
	"log"
	"net/http"
)

func Initialize() {
	server := http.NewServeMux()
	server.Handle("/", handlers.RootHandler())
	server.Handle("/message", handlers.Message())

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
