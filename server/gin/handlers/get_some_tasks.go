package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type GetSomeTasks struct{}

func (gst GetSomeTasks) Do(c *gin.Context) {
	page, ok := readPagination(c)
	if !ok {
		return
	}

	repo, ok := getTaskRepository(c)
	if !ok {
		return
	}

	ts, ok := gst.getSomeTaskFromDB(c, repo, page, c.GetInt("id"))
	if !ok {
		return
	}

	c.JSON(http.StatusOK, ts)
}

func (gst GetSomeTasks) getSomeTaskFromDB(c *gin.Context, repo database.TaskRepository, page, createdBy int) (ts []*task.Task, ok bool) {
	ts, err := repo.GetSomeTasks(page, createdBy)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
