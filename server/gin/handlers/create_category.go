package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type CreateCategory struct{}

func (cc CreateCategory) Do(c *gin.Context) {
	category, ok := cc.readCategory(c)
	if !ok {
		return
	}

	category, ok = cc.createCategory(c, category)
	if !ok {
		return
	}

	repo, ok := getCategoryRepository(c)
	if !ok {
		return
	}

	id, ok := cc.saveCategoryInDB(c, repo, category)
	if !ok {
		return
	}

	category.ID = id

	c.JSON(http.StatusCreated, category)
}

func (cc CreateCategory) readCategory(c *gin.Context) (category task.Category, ok bool) {
	ok = readRequestData(c, &category)
	return
}

func (cc CreateCategory) createCategory(c *gin.Context, cr task.Category) (category task.Category, ok bool) {
	if cr.CreatedBy == 0 {
		cr.CreatedBy = c.GetInt("id")
	}

	category, err := task.NewCategory(cr)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (cc CreateCategory) saveCategoryInDB(c *gin.Context, repo database.CategoryRepository, cr task.Category) (id int, ok bool) {
	id, err := repo.CreateCategory(cr)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
