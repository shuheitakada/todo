package data

import (
	"time"
)

type Task struct {
	ID          int
	Name        string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func Tasks() (tasks []Task, err error) {
	result := Db.Order("created_at asc").Find(&tasks)
	err = result.Error
	return
}

func (user *User) CreateTask(name string, description string) (err error) {
	task := Task{
		Name:        name,
		Description: description,
		UserID:      user.ID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	result := Db.Create(&task)
	err = result.Error
	return
}

func FindTaskById(id string) (task Task, err error) {
	result := Db.Find(&task, id)
	err = result.Error
	return
}

func (task *Task) Update() (err error) {
	task.UpdatedAt = time.Now()
	result := Db.Save(&task)
	err = result.Error
	return
}

func (task *Task) Delete() (err error) {
	result := Db.Delete(&task)
	err = result.Error
	return
}

func (task *Task) CreatedAtDate() string {
	return task.CreatedAt.Format("2006/1/2 15:04:05")
}

func (task *Task) UpdatedAtDate() string {
	return task.UpdatedAt.Format("2006/1/2 15:04:05")
}
