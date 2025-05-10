package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love Go!!!!")
}

func setHtmlResp(message []byte, httpCode int, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func main() {
	fmt.Println("Server listening on port 8080...")

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		tmpl, err := template.ParseFiles("index.tmpl", "button.tmpl")

		if err != nil {
			http.Error(res, "Failed to load template", http.StatusInternalServerError)
			fmt.Println("Error parsing template:", err)
			return
		}
		tmpl.ExecuteTemplate(res, "index.tmpl", nil)
		return
		
	})

	http.ListenAndServe(":8080", nil)
}
