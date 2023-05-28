package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/list"
	"github.com/gin-gonic/gin"
)

type GetSomeLists struct{}

func (gst GetSomeLists) Do(c *gin.Context) {
	page, ok := readPagination(c)
	if !ok {
		return
	}

	repo, ok := getListRepository(c)
	if !ok {
		return
	}

	ls, ok := gst.getSomeListFromDB(c, repo, page, c.GetInt("id"))
	if !ok {
		return
	}

	c.JSON(http.StatusOK, ls)
}

func (gst GetSomeLists) getSomeListFromDB(c *gin.Context, repo database.ListRepository, page, createdBy int) (ls []*list.List, ok bool) {
	ls, err := repo.GetSomeLists(page, createdBy)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
