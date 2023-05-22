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
	ok = l.checkCredentials(c, acct)
	if !ok {
		return
	}

	acct, ok = l.encryptCredentials(c, acct)
	if !ok {
		return
	}

	// Continue the search login at the database process
	_, ok = l.getAccountRepository(c)
	if !ok {
		return
	}
}

func (l Login) readCredentials(c *gin.Context) (acct account.Account, ok bool) {
	err := c.ShouldBind(&acct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ok = true
	return
}

func (l Login) checkCredentials(c *gin.Context, acct account.Account) (ok bool) {
	err := account.ValidateCredentials(acct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ok = true
	return
}

func (l Login) encryptCredentials(c *gin.Context, acct account.Account) (resAcct account.Account, ok bool) {
	p, err := utils.HashPassword(acct.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	resAcct.Password = p
	ok = true
	return
}

func (l Login) getAccountRepository(c *gin.Context) (repo database.AccountRepository, ok bool) {
	repo, err := database.GetAccountRepository(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ok = true
	return
}

func (l Login) searchCredentialsInDB(c *gin.Context, acct account.Account, repo database.AccountRepository) (id int, ok bool) {
	id, err := repo.MatchCredentials(acct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	return
}
