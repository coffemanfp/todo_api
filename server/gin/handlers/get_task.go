package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type GetTask struct{}

func (gt GetTask) Do(c *gin.Context) {
	id, ok := gt.readTaskID(c)
	if !ok {
		return
	}

	repo, ok := getTaskRepository(c)
	if !ok {
		return
	}

	t, ok := gt.getTaskFromDB(c, repo, id)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, t)
}

func (gt GetTask) readTaskID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (gt GetTask) getTaskFromDB(c *gin.Context, repo database.TaskRepository, id int) (t task.Task, ok bool) {
	t, err := repo.GetTask(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
