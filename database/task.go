package database

import (
	"github.com/coffemanfp/todo/search"
	"github.com/coffemanfp/todo/task"
)

const TASK_REPOSITORY RepositoryID = "TASK"

type TaskRepository interface {
	CreateTask(task task.Task) (id int, err error)
	UpdateTask(task task.Task) (err error)
	DeleteTask(id int) (err error)
	GetSomeTasks(page, createdBy int) (ts []*task.Task, err error)
	GetTask(id int) (task task.Task, err error)
	Search(search search.Search) (ts []*task.Task, err error)
}
