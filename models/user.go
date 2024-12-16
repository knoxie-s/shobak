package models

type User struct {
	ID       int64  `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}
