package handlers

import (
	"log"
	"net/http"
)

func Message() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello HTMX"))
		if err != nil {
			log.Printf("error writing message: %v", err)
		}
	})
}
