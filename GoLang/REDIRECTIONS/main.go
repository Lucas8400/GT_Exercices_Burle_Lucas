package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl1 := template.Must(template.ParseFiles("page1.html"))
	tmpl2 := template.Must(template.ParseFiles("page2.html"))
	tmpl3 := template.Must(template.ParseFiles("page3.html"))

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})
	http.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		tmpl1.Execute(w, nil)
	})
	http.HandleFunc("/page2", func(w http.ResponseWriter, r *http.Request) {
		tmpl2.Execute(w, nil)
	})
	http.HandleFunc("/page3", func(w http.ResponseWriter, r *http.Request) {
		tmpl3.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}
