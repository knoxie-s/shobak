package routes

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserReq struct {
	Login    string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

type UserFilter struct {
	Age   *int   `form:"age"`
	Login string `form:"login"`
	Email string `form:"email"`
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

func GetUserByID(c *gin.Context) {
	var user User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid url params!"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": cache[user.ID]})
}

func GetUser(c *gin.Context) {
	var userFilter UserFilter

	if err := c.Bind(&userFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid url params!"})
	}

	var users []User

	if userFilter.Login != "" {
		for _, v := range cache {
			if strings.Contains(v.Login, userFilter.Login) {
				users = append(users, v)
			}
		}
	}

	if userFilter.Email != "" {
		for _, v := range cache {
			if strings.Contains(v.Email, userFilter.Email) {
				users = append(users, v)
			}
		}
	}

	if userFilter.Age != nil {
		for _, v := range cache {
			if v.Age == *userFilter.Age {
				users = append(users, v)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
}
