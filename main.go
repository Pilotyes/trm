package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func login(login, password string) error {
	fmt.Println(login, password)
	if login == "a" && password == "b" {
		return nil
	}
	return errors.New("User not found")
}
func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles("index.html"))

		switch r.Method {
		case http.MethodGet:
			templates.ExecuteTemplate(rw, "index.html", nil)
		case http.MethodPost:
			if err := login(r.FormValue("login"), r.FormValue("password")); err == nil {
				http.Redirect(rw, r, "/lk", http.StatusMovedPermanently)
				return
			}
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(rw, "<html>Invalid password <a href=\"/login\">Back</a></html>")
		}

	})

	http.ListenAndServe(":8080", nil)
}
