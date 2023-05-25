package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

type Register struct{}

func (r Register) Do(c *gin.Context) {
	acct, ok := r.readAccount(c)
	if !ok {
		return
	}

	acct, ok = r.createNewAccount(c, acct)
	if !ok {
		return
	}

	repo, ok := r.getAccountRepository(c)
	if !ok {
		return
	}

	id, ok := r.registerAccount(c, acct, repo)
	if !ok {
		return
	}

	c.JSON(http.StatusCreated, id)

}

func (r Register) readAccount(c *gin.Context) (acct account.Account, ok bool) {
	return readAccount(c)
}

func (r Register) createNewAccount(c *gin.Context, acctR account.Account) (acct account.Account, ok bool) {
	acct, err := account.New(acctR)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ok = true
	return
}

func (r Register) getAccountRepository(c *gin.Context) (repo database.AccountRepository, ok bool) {
	return getAccountRepository(c)
}

func (r Register) registerAccount(c *gin.Context, acct account.Account, repo database.AccountRepository) (id int, ok bool) {
	id, err := repo.Register(acct)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ok = true
	return
}
