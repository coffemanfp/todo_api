package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

type DeleteCategory struct{}

func (dc DeleteCategory) Do(c *gin.Context) {
	id, ok := dc.readCategoryID(c)
	if !ok {
		return
	}

	repo, ok := getCategoryRepository(c)
	if !ok {
		return
	}

	ok = dc.deleteCategoryInDB(c, repo, id)
	if !ok {
		return
	}

	c.Status(http.StatusOK)
}

func (dc DeleteCategory) readCategoryID(c *gin.Context) (id int, ok bool) {
	return readIntFromURL(c, "id", false)
}

func (dc DeleteCategory) deleteCategoryInDB(c *gin.Context, repo database.CategoryRepository, id int) (ok bool) {
	err := repo.DeleteCategory(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
