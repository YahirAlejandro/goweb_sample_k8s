package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 301)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":80", nil)
}
