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
	server.ListenAndServe()
}
