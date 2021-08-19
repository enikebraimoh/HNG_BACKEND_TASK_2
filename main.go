package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/preocess", process)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, nil)
}

type Name struct {
	FName, LName string
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}

	tpl.ExecuteTemplate(w, "process.gohtml", d)

}
