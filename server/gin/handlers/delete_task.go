package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

type DeleteTask struct{}

func (dt DeleteTask) Do(c *gin.Context) {
	id, ok := dt.readTaskID(c)
	if !ok {
		return
	}

	repo, ok := getTaskRepository(c)
	if !ok {
		return
	}

	ok = dt.deleteTaskInDB(c, repo, id)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (dt DeleteTask) readTaskID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (dt DeleteTask) deleteTaskInDB(c *gin.Context, repo database.TaskRepository, id int) (ok bool) {
	err := repo.DeleteTask(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
