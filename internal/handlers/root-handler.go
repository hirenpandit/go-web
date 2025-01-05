package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func RootHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Root handler called\n")
		w.WriteHeader(http.StatusOK)
		temp, err := template.ParseFiles("templates/index.html")
		if err != nil {
			log.Printf("Error parsign template fiel %s : %v", "tempalte/index.html", err)
		}
		res := map[string]interface{}{
			"Name": "Hiren",
		}
		err = temp.Execute(w, res)
		if err != nil {
			log.Printf("Error parsign template fiel %s : %v", "templates/index.html", err)
		}
	})
}
