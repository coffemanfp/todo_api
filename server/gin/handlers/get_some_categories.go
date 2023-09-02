package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type GetSomeCategories struct{}

func (gst GetSomeCategories) Do(c *gin.Context) {
	page, ok := readPagination(c)
	if !ok {
		return
	}

	repo, ok := getCategoryRepository(c)
	if !ok {
		return
	}

	ts, ok := gst.getSomeCategoryFromDB(c, repo, page, c.GetInt("id"))
	if !ok {
		return
	}

	c.JSON(http.StatusOK, ts)
}

func (gst GetSomeCategories) getSomeCategoryFromDB(c *gin.Context, repo database.CategoryRepository, page, createdBy int) (ts []*task.Category, ok bool) {
	ts, err := repo.GetSomeCategories(page, createdBy)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
