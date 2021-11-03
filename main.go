package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("templates/*.html"))
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler).Methods("GET")
	mux.HandleFunc("/test", handler)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	templates.ExecuteTemplate(w, "index.html", "Hello my friend tonnytg")
}
