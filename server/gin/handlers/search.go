package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/search"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type Search struct{}

func (s Search) Do(c *gin.Context) {
	srch, ok := s.readSearch(c)
	if !ok {
		return
	}

	repo, ok := getTaskRepository(c)
	if !ok {
		return
	}

	ts, ok := s.searchOnDB(c, repo, srch)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, ts)
}

func (s Search) readSearch(c *gin.Context) (srch search.Search, ok bool) {
	title := c.Query("title")
	clientID := c.GetInt("id")

	isDone, ok := readBoolFromURL(c, "isDone", true)
	if !ok {
		return
	}
	isAddedToMyDay, ok := readBoolFromURL(c, "isAddedToMyDay", true)
	if !ok {
		return
	}
	isImportant, ok := readBoolFromURL(c, "isImportant", true)
	if !ok {
		return
	}
	hasDueDate, ok := readBoolFromURL(c, "hasDueDate", true)
	if !ok {
		return
	}
	expireSoon, ok := readBoolFromURL(c, "expireSoon", true)
	if !ok {
		return
	}

	srch, err := search.New(clientID, &title, isDone, isAddedToMyDay, isImportant, hasDueDate, expireSoon)
	if err != nil {
		handleError(c, err)
		return
	}

	ok = true
	return
}

func (s Search) searchOnDB(c *gin.Context, repo database.TaskRepository, srch search.Search) (ts []*task.Task, ok bool) {
	ts, err := repo.Search(srch)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
