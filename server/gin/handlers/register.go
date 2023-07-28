package handlers

import (
	"net/http"

	"github.com/coffemanfp/todo/account"
	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server/errors"
	"github.com/coffemanfp/todo/utils"
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

	id, ok := r.registerAccountInDB(c, acct, repo)
	if !ok {
		return
	}

	token, ok := r.generateToken(c, id)
	if !ok {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

}

func (r Register) readAccount(c *gin.Context) (acct account.Account, ok bool) {
	ok = readRequestData(c, &acct)
	return
}

func (r Register) createNewAccount(c *gin.Context, acctR account.Account) (acct account.Account, ok bool) {
	acct, err := account.New(acctR)
	if err != nil {
		err = errors.NewHTTPError(http.StatusBadRequest, err.Error())
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (r Register) getAccountRepository(c *gin.Context) (repo database.AuthRepository, ok bool) {
	return getAccountRepository(c)
}

func (r Register) registerAccountInDB(c *gin.Context, acct account.Account, repo database.AuthRepository) (id int, ok bool) {
	id, err := repo.Register(acct)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}

func (r Register) generateToken(c *gin.Context, id int) (token string, ok bool) {
	token, err := utils.GenerateToken(id, conf.Server.JWTLifespan, conf.Server.SecretKey)
	if err != nil {
		handleError(c, err)
		return
	}
	ok = true
	return
}
