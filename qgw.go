package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Will wait...")
	time.Sleep(5 * time.Second)
	fmt.Println("Moving on to login...")
	time.Sleep(1 * time.Second)
	http.Redirect(w, r, "/login", 301)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Landed on login...")
	time.Sleep(1 * time.Second)
	fmt.Println("Method: ", r.Method)
	time.Sleep(5 * time.Second)
	if r.Method == "GET" {
		t, err := template.ParseFiles("login.gtpl")
		//catch errors when parsing template
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])
	}
}

/*func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":80", nil)
}
*/
