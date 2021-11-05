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
	mux.HandleFunc("/", indexHandler).Methods("GET")
	mux.HandleFunc("/test", handler)
	mux.HandleFunc("/form", formHandler)

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

func indexHandler(w http.ResponseWriter, r *http.Request) {

	p := Person{
		"Jon",
		"Snow",
	}
	templates.ExecuteTemplate(w, "index.html", p)
}

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	if r.Method == http.MethodGet {
		templates.ExecuteTemplate(w, "forms.html", nil)
		return
	}

	templates.ExecuteTemplate(w, "forms.html", details)
	fmt.Println(details.Email)
	fmt.Println(details.Subject)
	fmt.Println(details.Message)

}
