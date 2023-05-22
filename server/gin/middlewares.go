package gin

import (
	"time"

	"github.com/coffemanfp/todo/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newCors(conf config.ConfigInfo) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     conf.Server.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
