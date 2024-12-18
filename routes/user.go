package routes

import (
	"fmt"
	"log"
	"net/http"
	"shobak/db"
	"shobak/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserReq struct {
	Login    string `json:"name"`
	Email    string `json:"email"`
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

var cache = make(map[int64]models.User, 1000) // "1": User{}

func CreateUser(c *gin.Context) {
	var userReq models.User

	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid body request!"})
		return
	}

	user := models.User{
		Login:    userReq.Login,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	var count int64
	db.GetDB().Model(models.User{}).Where("login = ?", userReq.Login).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": fmt.Sprintf("user %s already exist", userReq.Login)})
		return
	}

	if err := db.GetDB().Create(&user).Error; err != nil {
		log.Printf("error to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": UserResp{
		ID: user.ID,
	}})
}

func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	var user models.User

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("error to parse id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
	}

	if err = db.GetDB().Where("id = ?", id).Find(&user).Error; err != nil {
		log.Printf("error to get user by id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "internal server error"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}

func GetUser(c *gin.Context) {
	var userFilter UserFilter

	if err := c.Bind(&userFilter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "provide valid url params!"})
	}

	var users []models.User

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
