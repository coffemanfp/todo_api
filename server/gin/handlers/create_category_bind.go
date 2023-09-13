package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/task"
	"github.com/gin-gonic/gin"
)

type CreateCategoryBind struct{}

func (ccb CreateCategoryBind) Do(c *gin.Context) {
	binds, ok := ccb.readCategoryBinds(c)
	if !ok {
		return
	}

	repo, ok := getCategoryRepository(c)
	if !ok {
		return
	}

	ok = ccb.saveCategoryTaskBindsInDB(c, repo, binds)
	if !ok {
		return
	}

	c.Status(http.StatusCreated)
}

func (ccb CreateCategoryBind) readCategoryBinds(c *gin.Context) (binds []*task.CategoryTaskBind, ok bool) {
	ok = readRequestData(c, &binds)
	return
}

func (ccb CreateCategoryBind) saveCategoryTaskBindsInDB(c *gin.Context, repo database.CategoryRepository, binds []*task.CategoryTaskBind) (ok bool) {
	err := repo.CreateCategoryBinds(binds)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
