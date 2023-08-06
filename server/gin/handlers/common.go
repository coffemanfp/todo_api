package handlers

import (
	"fmt"
	"strconv"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

func readRequestData(c *gin.Context, v interface{}) (ok bool) {
	err := c.ShouldBindJSON(v)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func validateCredentials(c *gin.Context, acct account.Account) (ok bool) {
	err := account.ValidateCredentials(acct)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func getAccountRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	repo, err := database.GetRepository[database.AuthRepository](db, database.AUTH_REPOSITORY)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func getListRepository(c *gin.Context) (repo database.ListRepository, ok bool) {
	repo, err := database.GetRepository[database.ListRepository](db, database.LIST_REPOSITORY)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func getTaskRepository(c *gin.Context) (repo database.TaskRepository, ok bool) {
	repo, err := database.GetRepository[database.TaskRepository](db, database.TASK_REPOSITORY)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func readIntParam(c *gin.Context, param string) (v int, ok bool) {
	p := c.Param(param)
	if p == "" {
		ok = true
		return
	}
	v, err := strconv.Atoi(p)
	if err != nil {
		err = fmt.Errorf("invalid %s param: %s", param, p)
		handleError(c, err)
		return
	}

	ok = true
	return
}

func readPagination(c *gin.Context) (page int, ok bool) {
	page, ok = readIntParam(c, "page")
	return
}

func handleError(c *gin.Context, err error) {
	c.Error(err)
	c.Abort()
}
