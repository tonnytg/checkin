package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("templates/*.html"))
	mux := mux.NewRouter()
	mux.HandleFunc("/", formHandler)

	http.ListenAndServe(":8081", mux)

	fmt.Println("Server started")

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type Person struct {
	Name     string
	Lastname string
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "forms.html", nil)
		return
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	templates.ExecuteTemplate(w, "index.html", details)
}
