package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", process)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "80" // Default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, nil)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("process.html")
		t.Execute(w, nil)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	d := struct {
		Name    string
		Email   string
		Message string
	}{
		Name:    name,
		Email:   email,
		Message: message,
	}

	tpl.ExecuteTemplate(w, "process.html", d)

}
