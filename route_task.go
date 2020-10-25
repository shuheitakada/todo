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
		"templates/public.navbar.html",
		"templates/tasks/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", tasks)
}

func handleNewTask(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
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

func handleEditTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, err := data.FindTaskById(id)
	if err != nil {
		return
	}
	files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/tasks/edit.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", task)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	// formからname属性を取得
	err := r.ParseForm()
	if err != nil {
		return
	}
	id := r.PostFormValue("task_id")
	name := r.PostFormValue("task_name")
	description := r.PostFormValue("task_description")
	task, err := data.FindTaskById(id)
	if err != nil {
		return
	}

	// taskを更新
	task.Name = name
	task.Description = description
	task.Update()

	http.Redirect(w, r, "/", 302)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, err := data.FindTaskById(id)
	if err != nil {
		return
	}
	task.Delete()

	http.Redirect(w, r, "/", 302)
}
