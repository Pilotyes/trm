package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"trm/internal/store"
)

func login(login, password string) error {
	if u := store.FindUser(login); u != nil {
		if u.Password == password {
			return nil
		} else {
			return errors.New("Incorrect password")
		}
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
			err := login(r.FormValue("login"), r.FormValue("password"))
			if err != nil {
				rw.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(rw, "<html>%s <a href=\"/login\">Back</a></html>", err)
				return
			}
			http.Redirect(rw, r, "/lk", http.StatusMovedPermanently)
		}

	})

	http.ListenAndServe(":8080", nil)
}
