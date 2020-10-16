package main

import (
	"html/template"
	"net/http"
	"todo/data"
)

func handleTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := data.Tasks()
	if err != nil {
		return
	}
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/tasks/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", tasks)
}

func handleNewTask(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/tasks/new.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", nil)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	name := r.PostFormValue("task_name")
	description := r.PostFormValue("task_description")
	_, err = data.CreateTask(name, description)
	if err != nil {
		return
	}
	http.Redirect(w, r, "/", 302)
}
