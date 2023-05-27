package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/list"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/gin-gonic/gin"
)

type UpdateList struct{}

func (ul UpdateList) Do(c *gin.Context) {
	l, ok := ul.readList(c)
	if !ok {
		return
	}

	id, ok := ul.readListID(c)
	if !ok {
		return
	}

	l, ok = ul.updateList(c, id, l)
	if !ok {
		return
	}

	repo, ok := getListRepository(c)
	if !ok {
		return
	}

	ok = ul.updateListInDB(c, repo, l)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (ul UpdateList) readListID(c *gin.Context) (id int, ok bool) {
	return readIntParam(c, "id")
}

func (ul UpdateList) readList(c *gin.Context) (l list.List, ok bool) {
	ok = readRequestData(c, &l)
	return
}

func (ul UpdateList) updateList(c *gin.Context, id int, lR list.List) (l list.List, ok bool) {
	l, err := list.Update(lR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	l.ID = id
	ok = true
	return
}

func (ul UpdateList) updateListInDB(c *gin.Context, repo database.ListRepository, l list.List) (ok bool) {
	err := repo.UpdateList(l)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
