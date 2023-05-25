package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

func readAccount(c *gin.Context) (acct account.Account, ok bool) {
	err := c.ShouldBind(&acct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ok = true
	return
}

func validateCredentials(c *gin.Context, acct account.Account) (ok bool) {
	err := account.ValidateCredentials(acct)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ok = true
	return
}

func getAccountRepository(c *gin.Context) (repo database.AccountRepository, ok bool) {
	repo, err := database.GetAccountRepository(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ok = true
	return
}
