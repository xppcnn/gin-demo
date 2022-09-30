package system

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/xppcnn/gin-demo/utils"
)

type UserService struct {
}

type User struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	UserName string    `json:"user_name"`
	PassWord string    `json:"pass_word"`
}

type PublicUser struct {
	*User              // 匿名嵌套
	Password *struct{} `json:"pass_word,omitempty"`
}

func (userService *UserService) Register(u User) (userInfo PublicUser, err error) {
	var vo User
	if !errors.Is(db.Where("user_name = ?", u.UserName).First(&vo).Error, gorm.ErrRecordNotFound) {
		return userInfo, errors.New("用户已注册")
	}
	u.PassWord = utils.BcryptHash(u.PassWord)
	u.ID = uuid.New()
	err = db.Create(u).Error
	pu := PublicUser{User: &u}
	return pu, err
}

func (userService *UserService) FindUserById(id string) (user *PublicUser, err error) {
	var u User
	if err = db.Where("`id` = ?", id).First(&u).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	user = &PublicUser{User: &u}
	return user, nil
}
