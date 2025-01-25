package server

import (
	"go-web/internal/handlers"
	"go-web/internal/handlers/ai"
	"go-web/internal/server/middleware"
	"go-web/internal/websocket"
	"log"
	"net/http"
)

func Initialize() {
	server := http.NewServeMux()
	server.Handle("/", middleware.Logger(handlers.RootHandler()))
	server.Handle("/message", middleware.Logger(handlers.Message()))

	//AI chat related handlers
	server.Handle("/chat", middleware.Logger(ai.ChatHandler()))
	server.Handle("/ai/query", middleware.Logger(ai.QueryHandler()))

	//websocket
	server.Handle("/ws", websocket.Handler())

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalf("Error initializing server: %v", err)
	}
}
