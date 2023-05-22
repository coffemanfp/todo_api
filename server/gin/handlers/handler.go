package handlers

import "github.com/gin-gonic/gin"

type Handler interface {
	Do(c *gin.Context)
}
