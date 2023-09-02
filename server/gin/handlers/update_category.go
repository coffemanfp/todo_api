package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type UpdateCategory struct{}

func (uc UpdateCategory) Do(c *gin.Context) {
	category, ok := uc.readCategory(c)
	if !ok {
		return
	}

	id, ok := uc.readCategoryID(c)
	if !ok {
		return
	}

	repo, ok := getCategoryRepository(c)
	if !ok {
		return
	}

	category.ID = id
	ok = uc.updateCategoryInDB(c, repo, category)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (uc UpdateCategory) readCategory(c *gin.Context) (category task.Category, ok bool) {
	ok = readRequestData(c, &category)
	return
}

func (uc UpdateCategory) readCategoryID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (cc CreateCategory) updateCategory(c *gin.Context, cr task.Category) (category task.Category, ok bool) {
	category, err := task.UpdateCategory(cr)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (uc UpdateCategory) updateCategoryInDB(c *gin.Context, repo database.CategoryRepository, category task.Category) (ok bool) {
	err := repo.UpdateCategory(category)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
