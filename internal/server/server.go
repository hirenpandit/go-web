package server

import (
	"go-web/internal/handlers"
	"go-web/internal/server/middleware"
	"log"
	"net/http"
)

func Initialize() {
	server := http.NewServeMux()
	server.Handle("/", middleware.Logger(handlers.RootHandler()))
	server.Handle("/message", middleware.Logger(handlers.Message()))
	server.Handle("/chat", middleware.Logger(handlers.ChatHandler()))

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
