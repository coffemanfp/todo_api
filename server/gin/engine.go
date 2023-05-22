package gin

import (
	"github.com/coffemanfp/todo/config"
	"github.com/coffemanfp/todo/database"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	conf config.ConfigInfo
	db   database.Database
	r    *gin.Engine
}

func (ge GinEngine) New(conf config.Config, db database.Database) {
	ge.conf = conf.Get()
	ge.r = gin.Default()
	// setup the middlewares
}

func (ge GinEngine) setAccountHandlers() {
	_, err := database.GetAccountRepository(ge.db.Repositories)
	if err != nil {
		return
	}

	// set up the login handler
}
