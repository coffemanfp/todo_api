package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

type DeleteList struct{}

func (dl DeleteList) Do(c *gin.Context) {
	id, ok := dl.readListId(c)
	if !ok {
		return
	}

	repo, ok := getListRepository(c)
	if !ok {
		return
	}

	ok = dl.deleteListFromDB(c, repo, id)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (dl DeleteList) readListId(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (dl DeleteList) deleteListFromDB(c *gin.Context, repo database.ListRepository, id int) (ok bool) {
	err := repo.DeleteList(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
