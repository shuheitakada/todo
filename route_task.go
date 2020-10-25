package main

import (
	"net/http"
	"todo/data"
)

func handleTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := data.Tasks()
	if err != nil {
		return
	}
	generateHTML(w, tasks, "layout", "public.navbar", "tasks/index")
}

func handleNewTask(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "layout", "public.navbar", "tasks/new")
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
	generateHTML(w, task, "layout", "public.navbar", "tasks/edit")
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
