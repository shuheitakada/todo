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
	rows, err := Db.Query("SELECT id, name, description, created_at, updated_at FROM tasks ORDER BY id ASC")
	if err != nil {
		return
	}
	for rows.Next() {
		task := Task{}
		err = rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return
		}
		tasks = append(tasks, task)
	}
	rows.Close()
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
