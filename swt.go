package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func new_entry(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("new_entry.html")
	t.Execute(w, nil)
}

func submit_entry(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		entry := r.Form["entry"]
		fmt.Fprintf(w, "This is the post request: ")
		fmt.Fprint(w, entry)
	} else {
		http.Redirect(w, r, "/new-entry", 301)
	}
}

func main() {
	http.HandleFunc("/new-entry", new_entry)
	http.ListenAndServe(":80", nil)
}
