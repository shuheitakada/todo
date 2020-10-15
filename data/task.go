package data

import (
	"time"
)

type Task struct {
	ID           int
	Name         string
	Descriprtion string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func CreateTask(name string, description string) (task Task, err error) {
	task = Task{}
	err = Db.QueryRow("INSERT INTO tasks (name, description, created_at, updated_at) values ($1, $2, $3, $4) returning id, name, description, created_at, updated_at", name, description, time.Now(), time.Now()).Scan(&task.ID, &task.Name, &task.Descriprtion, &task.CreatedAt, &task.UpdatedAt)
	return
}
