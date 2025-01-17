package handlers

import (
	"go-web/internal/anthropic"
	"html/template"
	"log"
	"net/http"
)

func ChatHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		temp, err := template.ParseFiles("templates/anthropic-chat.html")
		if err != nil {
			log.Printf("error parsing template anthropic-chat.html: %v", err)
		}
		res := make(map[string]interface{})
		res["Message"] = anthropic.SendMessage()
		err = temp.Execute(w, res)
		if err != nil {
			log.Printf("error executing template anthropic-chat.html: %v", err)
		}
	})
}
