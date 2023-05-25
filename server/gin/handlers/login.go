package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/utils"
	"github.com/gin-gonic/gin"
)

type Login struct{}

func (l Login) Do(c *gin.Context) {
	acct, ok := l.readCredentials(c)
	if !ok {
		return
	}
	ok = l.validateCredentials(c, acct)
	if !ok {
		return
	}

	acct, ok = l.encryptCredentials(c, acct)
	if !ok {
		return
	}

	// Continue the search login at the database process
	repo, ok := l.getAccountRepository(c)
	if !ok {
		return
	}

	id, ok := l.searchCredentialsInDB(c, acct, repo)
	if !ok {
		return
	}

	token, ok := l.generateToken(c, id)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (l Login) readCredentials(c *gin.Context) (acct account.Account, ok bool) {
	return readAccount(c)
}

func (l Login) validateCredentials(c *gin.Context, acct account.Account) (ok bool) {
	return validateCredentials(c, acct)
}

func (l Login) encryptCredentials(c *gin.Context, acct account.Account) (resAcct account.Account, ok bool) {
	p, err := utils.HashPassword(acct.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	resAcct.Password = p
	ok = true
	return
}

func (l Login) getAccountRepository(c *gin.Context) (repo database.AccountRepository, ok bool) {
	return getAccountRepository(c)
}

func (l Login) searchCredentialsInDB(c *gin.Context, acct account.Account, repo database.AccountRepository) (id int, ok bool) {
	id, err := repo.MatchCredentials(acct)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ok = true
	return
}

func (l Login) generateToken(c *gin.Context, id int) (token string, ok bool) {
	token, err := utils.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	ok = true
	return
}
