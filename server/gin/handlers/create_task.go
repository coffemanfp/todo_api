package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type CreateTask struct{}

func (ct CreateTask) Do(c *gin.Context) {
	t, ok := ct.readTask(c)
	if !ok {
		return
	}

	t, ok = ct.createTask(c, t)
	if !ok {
		return
	}

	repo, ok := getTaskRepository(c)
	if !ok {
		return
	}

	id, ok := ct.saveTaskInDB(c, repo, t)
	if !ok {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (ct CreateTask) readTask(c *gin.Context) (t task.Task, ok bool) {
	ok = readRequestData(c, &t)
	return
}

func (ct CreateTask) createTask(c *gin.Context, tR task.Task) (t task.Task, ok bool) {
	if tR.CreatedBy == 0 {
		tR.CreatedBy = c.GetInt("id")
	}

	t, err := task.New(tR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (ct CreateTask) saveTaskInDB(c *gin.Context, repo database.TaskRepository, t task.Task) (id int, ok bool) {
	id, err := repo.CreateTask(t)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
