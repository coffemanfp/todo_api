package gin

import (
	"github.com/coffemanfp/todo/config"
	"github.com/coffemanfp/todo/database"
	"github.com/coffemanfp/todo/server"
	"github.com/coffemanfp/todo/server/gin/handlers"
	"github.com/gin-gonic/gin"
)

type GinEngine struct {
	conf config.ConfigInfo
	db   database.Database
	r    *gin.Engine
}

func New(conf config.ConfigInfo, db database.Database) server.Engine {
	ge := GinEngine{
		conf: conf,
		db:   db,
		r:    gin.New(),
	}

	handlers.Init(ge.db.Repositories, ge.conf)

	v1 := ge.r.Group("/v1")

	ge.setCommonMiddlewares(v1)
	ge.setAccountHandlers(v1)
	return ge.r
}

func (ge GinEngine) setAccountHandlers(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", handlers.Login{}.Do)
	auth.POST("/register", handlers.Register{}.Do)
}

func (ge GinEngine) setCommonMiddlewares(r *gin.RouterGroup) {
	r.Use(newCors(ge.conf))
	r.Use(gin.Recovery())
	r.Use(logger())
}
