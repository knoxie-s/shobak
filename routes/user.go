package routes

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserReq struct {
	Login    string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int64  `json:"id" form:"id"`
	Login    string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResp struct {
	ID int64 `json:"id"`
}

var cache = make(map[int64]User, 1000) // "1": User{}

func CreateUser(c *gin.Context) {
	var userReq User

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid body request!"})
		return
	}

	id := rand.Int63n(1000)

	user := User{
		ID:       id,
		Login:    userReq.Login,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	cache[id] = user

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": UserResp{
		ID: id,
	}})
}

func GetUser(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid url params!"})
	}

	//v, ok := cache[user.ID]

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": cache[user.ID]})
}
