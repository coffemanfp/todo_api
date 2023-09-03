package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/gin-gonic/gin"
)

type CreateCategoryBind struct{}

func (ccb CreateCategoryBind) Do(c *gin.Context) {
	taskID, categoryID, ok := ccb.readCategoryBind(c)
	if !ok {
		return
	}

	repo, ok := getCategoryRepository(c)
	if !ok {
		return
	}

	id, ok := ccb.saveCategoryInDB(c, repo, taskID, categoryID)
	if !ok {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (ccb CreateCategoryBind) readCategoryBind(c *gin.Context) (taskID, categoryID int, ok bool) {
	taskID, ok = readIntFromURL(c, "taskId", true)
	if !ok {
		return
	}
	categoryID, ok = readIntFromURL(c, "categoryId", true)
	if !ok {
		return
	}

	if taskID == 0 {
		err := errors.NewHTTPError(http.StatusUnprocessableEntity, "invalid task id: task id can't be equals or lower than zero (0)")
		handleError(c, err)
		ok = false
		return
	}

	if categoryID == 0 {
		err := errors.NewHTTPError(http.StatusUnprocessableEntity, "invalid category id: category id can't be equals or lower than zero (0)")
		handleError(c, err)
		ok = false
	}
	return
}

func (ccb CreateCategoryBind) saveCategoryInDB(c *gin.Context, repo database.CategoryRepository, taskID, categoryID int) (id int, ok bool) {
	id, err := repo.CreateCategoryBind(taskID, categoryID)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
