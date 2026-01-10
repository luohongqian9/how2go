package model

import "ginlearn/global"

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserByID(id string) (User, error) {
	var user User
	err := global.DB.First(&user, id).Error
	return user, err
}
