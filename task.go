package main

import (
	"html/template"
	"net/http"
)

func handleTasks(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/tasks/index.html",
		"templates/navbar.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", nil)
}
