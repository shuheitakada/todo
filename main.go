package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started")

	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	http.HandleFunc("/", handleTasks)
	http.HandleFunc("/tasks/new", handleNewTask)
	http.HandleFunc("/tasks/create", handleCreateTask)
	http.HandleFunc("/tasks/edit", handleEditTask)
	http.HandleFunc("/tasks/update", handleUpdateTask)
	http.HandleFunc("/tasks/delete", handleDeleteTask)

	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/signup", handleSignup)

	http.HandleFunc("/favicon.ico", doNothing)
	server.ListenAndServe()
}
