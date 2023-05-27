package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
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

	// Continue the search login at the database process
	repo, ok := l.getAccountRepository(c)
	if !ok {
		return
	}

	id, hash, ok := l.searchCredentialsInDB(c, acct, repo)
	if !ok {
		return
	}

	ok = l.comparePassword(c, hash, acct.Password)
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
	ok = readRequestData(c, &acct)
	return
}

func (l Login) validateCredentials(c *gin.Context, acct account.Account) (ok bool) {
	return validateCredentials(c, acct)
}

func (l Login) getAccountRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	return getAccountRepository(c)
}

func (l Login) searchCredentialsInDB(c *gin.Context, acct account.Account, repo database.AuthRepository) (id int, hash string, ok bool) {
	id, hash, err := repo.GetIdAndHashedPassword(acct)
	if err != nil {
		err = errors.NewHTTPError(http.StatusUnauthorized, errors.UNAUTHORIZED_ERROR_MESSAGE)
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (l Login) comparePassword(c *gin.Context, hash, password string) (ok bool) {
	err := utils.CompareHashAndPassword(hash, password)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (l Login) generateToken(c *gin.Context, id int) (token string, ok bool) {
	token, err := utils.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
