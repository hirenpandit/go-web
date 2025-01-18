package ai

import (
	"go-web/internal/anthropic"
	"log"
	"net/http"
)

func QueryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		query := r.URL.Query().Get("message")
		if query == "" {
			log.Printf("query is empty/n")
		}
		anthropic.SendMessageWithStreamResponse(query)
		//response := anthropic.SendMessage(query)
		//_, err := w.Write([]byte(response))

		//if err != nil {
		//	log.Printf("failed to write response: %v", err)
		//}
	})
}
