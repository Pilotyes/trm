package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"trm/internal/store"
	"trm/internal/store/sessions"
)

const templatesDir = "./templates/"

var templatesList = []string{
	templatesDir + "index.html",
	templatesDir + "lk_co.html",
	templatesDir + "lk_en.html",
	templatesDir + "login.html",
}

func login(login, password string) error {
	if u := store.FindUser(login); u != nil {
		if u.Password == password {
			return nil
		}
		return errors.New("Incorrect password")
	}

	return errors.New("User not found")
}
func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles(templatesList...))

		switch r.Method {
		case http.MethodGet:
			templates.ExecuteTemplate(rw, "login.html", nil)
		case http.MethodPost:
			logUser, pasUser := r.FormValue("login"), r.FormValue("password")
			err := login(logUser, pasUser)
			if err != nil {
				rw.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(rw, "<html>%s <a href=\"/login\">Back</a></html>", err)
				return
			}
			uniqID := sessions.GetUniqSessionID()
			if uniqID == "" {
				rw.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(rw, "Please try again", err)
				return
			}
			c := http.Cookie{
				Name:  "SessionID",
				Value: uniqID,
			}
			http.SetCookie(rw, &c)
			sessions.Sessions[uniqID] = store.FindUser(logUser)
			http.Redirect(rw, r, "/lk", http.StatusMovedPermanently)
		}

	})

	http.HandleFunc("/lk", func(rw http.ResponseWriter, r *http.Request) {

		templates := template.Must(template.ParseFiles(templatesList...))
		cookieID, err := r.Cookie("SessionID")
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(rw, "No pasaran", err)
			return
		}
		user, ok := sessions.Sessions[cookieID.Value]
		if !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(rw, "No pasaran", err)
			return
		}
		switch user.UserType {
		case store.UserTypeM:
			templates.ExecuteTemplate(rw, "lk_en.html", nil)
		case store.UserTypeC:
			templates.ExecuteTemplate(rw, "lk_co.html", nil)
		default:
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(rw, "User type not found", err)
			//return
		}
	})
	http.ListenAndServe(":8080", nil)
}
