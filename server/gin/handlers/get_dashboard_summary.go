package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/dashboard"
	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

type GetDashboardSummary struct{}

func (gds GetDashboardSummary) Do(c *gin.Context) {
	repo, ok := getDashboardRepository(c)
	if !ok {
		return
	}

	id, ok := gds.readUserID(c)
	if !ok {
		return
	}

	info, ok := gds.getDashboardSummary(c, id, repo)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, info)
}

func (gds GetDashboardSummary) readUserID(c *gin.Context) (int, bool) {
	return readIntFromURL(c, "id", false)
}

func (gds GetDashboardSummary) getDashboardSummary(c *gin.Context, id int, repo database.DashboardRepository) (summary dashboard.DashboardSummary, ok bool) {
	summary, err := repo.GetDashboardSummary(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
