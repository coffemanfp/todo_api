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

	listID, ok := gst.readListID(c)
	if !ok {
		return
	}

	ts, ok := gst.getSomeTaskFromDB(c, repo, page, c.GetInt("id"), listID)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, ts)
}

func (gst GetSomeTasks) readListID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "listID", true)
}

func (gst GetSomeTasks) getSomeTaskFromDB(c *gin.Context, repo database.TaskRepository, page, createdBy, listID int) (ts []*task.Task, ok bool) {
	ts, err := repo.GetSomeTasks(page, listID, createdBy)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
