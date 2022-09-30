package models

import (
	"errors"

	"github.com/xppcnn/gin-demo/utils"
)

func CheckAccount(username, password string) (userInfo *PublicUser, err error) {
	var user User
	if err := db.Where("user_name = ?", username).First(&user).Error; err == nil {
		if ok := utils.BcryptCheck(password, user.PassWord); ok {
			pu := PublicUser{User: &user}
			return &pu, err
		} else {
			return nil, errors.New("密码错误")
		}
	}
	return nil, err
}
