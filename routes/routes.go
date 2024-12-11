package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	// CRUD - CREATE, READ, UPDATE, DELETE
	// Initialize default gin router
	defaultRouter := gin.Default()

	defaultRouter.GET("/ping", Ping)
	defaultRouter.POST("/users", CreateUser) // Создание пользователя
	defaultRouter.GET("/users", GetUser)     // Получение пользователя

	defaultRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return defaultRouter
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong..."})
}
