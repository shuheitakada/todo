package data

import (
	"time"
)

type Task struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Tasks() (tasks []Task, err error) {
	result := DbGorm.Order("created_at asc").Find(&tasks)
	err = result.Error
	return
}

func CreateTask(name string, description string) (task Task, err error) {
	task = Task{}
	err = Db.QueryRow("INSERT INTO tasks (name, description, created_at, updated_at) values ($1, $2, $3, $4) returning id, name, description, created_at, updated_at", name, description, time.Now(), time.Now()).Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	return
}

func FindTaskById(id int) (task Task, err error) {
	task = Task{}
	err = Db.QueryRow("SELECT id, name, description, created_at, updated_at FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.UpdatedAt)
	return
}

func (task *Task) Update() (err error) {
	_, err = Db.Exec("UPDATE tasks SET name = $2, description = $3, updated_at = $4 WHERE id = $1", task.ID, task.Name, task.Description, time.Now())
	return
}

func (task *Task) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM tasks WHERE id = $1", task.ID)
	return
}

func (task *Task) CreatedAtDate() string {
	return task.CreatedAt.Format("2006/1/2 15:04:05")
}

func (task *Task) UpdatedAtDate() string {
	return task.UpdatedAt.Format("2006/1/2 15:04:05")
}
