package service

import (
	"errors"
	"ginlearn/model"
)

func GetUserInfo(id string) (model.User, error) {
	user, err := model.GetUserByID(id)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}
