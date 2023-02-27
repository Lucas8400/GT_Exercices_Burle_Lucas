package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ChampSaisie struct {
	Nom     string
	Message string
}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		details := ChampSaisie{
			Nom:     r.FormValue("nom"),
			Message: r.FormValue("message"),
		}
		fmt.Println(details)
		tmpl.Execute(w, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
