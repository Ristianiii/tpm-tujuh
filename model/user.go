package model

type User struct {
	UserId   int    `json:"user_id" gorm:"column:user_id;autoIncrement"`
	Email    string `json:"email" gorm:"column:email;unique"`
	Password string `json:"password" gorm:"column:password"`
}
