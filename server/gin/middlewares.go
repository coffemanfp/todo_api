package gin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/coffemanfp/todo/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func logger() gin.HandlerFunc {
	return structuredLogger(&log.Logger)
}

func structuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}
		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg(param.ErrorMessage)
	}
}
func readToken(c *gin.Context) (token string, err error) {
	token = c.Query("token")
	if token != "" {
		return
	}

	token = c.Request.Header.Get("Authorization")
	if v := strings.Split(token, " "); len(v) == 2 {
		token = v[1]
	}
	if token == "" {
		err = errors.New("no token provided")
		c.AbortWithError(http.StatusUnauthorized, err)
	}
	return
}

func saveTokenContent(c *gin.Context, secretKey string) (err error) {
	tokenS, err := readToken(c)
	if err != nil {
		return
	}

	token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = errors.New("invalid token")
	}

	c.Set("id", claims["account_id"])
	return
}
