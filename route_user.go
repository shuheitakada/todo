package main

import (
	"html/template"
	"net/http"
	"todo/data"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/users/login.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", nil)
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		files := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/users/signup.html",
		}
		templates := template.Must(template.ParseFiles(files...))
		templates.ExecuteTemplate(w, "layout", nil)
	case "POST":
		if err := r.ParseForm(); err != nil {
			return
		}
		user := data.User{
			Email: r.PostFormValue("email"),
			CryptedPassword: r.PostFormValue("password"),
		}
		if err := user.Create(); err != nil {
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}
