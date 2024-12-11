package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	// Initialize default gin router
	defaultRouter := gin.Default()

	defaultRouter.GET("/ping", Ping)

	defaultRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return defaultRouter
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong..."})
}
