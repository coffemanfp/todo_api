package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type UpdateTask struct{}

func (ut UpdateTask) Do(c *gin.Context) {
	t, ok := ut.readTask(c)
	if !ok {
		return
	}

	id, ok := ut.readTaskID(c)
	if !ok {
		return
	}

	t, ok = ut.updateTask(c, id, t)
	if !ok {
		return
	}

	repo, ok := getTaskRepository(c)
	if !ok {
		return
	}

	ok = ut.updateTaskInDB(c, repo, t)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (ut UpdateTask) readTask(c *gin.Context) (t task.Task, ok bool) {
	ok = readRequestData(c, &t)
	return
}

func (ut UpdateTask) updateTask(c *gin.Context, id int, tR task.Task) (t task.Task, ok bool) {
	t, err := task.Update(tR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	t.ID = id
	ok = true
	return
}

func (ut UpdateTask) readTaskID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (ut UpdateTask) updateTaskInDB(c *gin.Context, repo database.TaskRepository, t task.Task) (ok bool) {
	err := repo.UpdateTask(t)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
