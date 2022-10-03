package system

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/xppcnn/gin-demo/global"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/utils"
)

type UserService struct {
}

func (userService *UserService) Register(u models.User) (userInfo models.PublicUser, err error) {
	var vo models.User
	if !errors.Is(global.Db.Where("user_name = ?", u.UserName).First(&vo).Error, gorm.ErrRecordNotFound) {
		return userInfo, errors.New("用户已注册")
	}
	u.PassWord = utils.BcryptHash(u.PassWord)
	u.ID = uuid.New()
	err = global.Db.Create(u).Error
	pu := models.PublicUser{User: &u}
	return pu, err
}

func (userService *UserService) FindUserById(id string) (user *models.PublicUser, err error) {
	var u models.User
	if err = global.Db.Where("`id` = ?", id).First(&u).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	user = &models.PublicUser{User: &u}
	return user, nil
}
