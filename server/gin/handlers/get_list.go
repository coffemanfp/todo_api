package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/list"
	"github.com/gin-gonic/gin"
)

type GetList struct{}

func (gl GetList) Do(c *gin.Context) {
	repo, ok := getListRepository(c)
	if !ok {
		return
	}

	id, ok := gl.readListID(c)
	if !ok {
		return
	}

	l, ok := gl.getListFromDB(c, id, repo)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, l)
}

func (gl GetList) readListID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (gl GetList) getListFromDB(c *gin.Context, id int, repo database.ListRepository) (l list.List, ok bool) {
	l, err := repo.GetList(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
