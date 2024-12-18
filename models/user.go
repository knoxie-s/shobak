package models

type User struct {
	ID       int64  `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	Login    string `json:"login" gorm:"column:login"`
	Email    string `json:"email" gorm:"email"`
	Age      int    `json:"age" gorm:"age"`
	Password string `json:"password" gorm:"password"`
}

func (User) TableName() string {
	return "users"
}
