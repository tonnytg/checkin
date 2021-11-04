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

	http.ListenAndServe(":8081", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

type Person struct {
    Name string
    Lastname  string
}

func indexHandler(w http.ResponseWriter, r *http.Request){

	p := Person{
		"Jon",
		"Snow",
	}
	templates.ExecuteTemplate(w, "index.html", p)
}
