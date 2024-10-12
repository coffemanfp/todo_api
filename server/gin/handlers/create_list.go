package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/list"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/gin-gonic/gin"
)

type CreateList struct{}

func (cl CreateList) Do(c *gin.Context) {
	l, ok := cl.readList(c)
	if !ok {
		return
	}

	l, ok = cl.createNewList(c, l)
	if !ok {
		return
	}

	repo, ok := getListRepository(c)
	if !ok {
		return
	}

	id, ok := cl.saveListInDB(c, repo, l)
	if !ok {
		return
	}

	l.ID = id

	c.JSON(http.StatusCreated, l)
}

func (cl CreateList) readList(c *gin.Context) (l list.List, ok bool) {
	ok = readRequestData(c, &l)
	return
}

func (cl CreateList) createNewList(c *gin.Context, lR list.List) (l list.List, ok bool) {
	if lR.CreatedBy == 0 {
		lR.CreatedBy = c.GetInt("id")
	}
	l, err := list.New(lR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (cl CreateList) saveListInDB(c *gin.Context, repo database.ListRepository, l list.List) (id int, ok bool) {
	id, err := repo.CreateList(l)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
