package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primary_key" json:"id"`
	UserName string    `json:"user_name"`
	PassWord string    `json:"pass_word"`
}

type PublicUser struct {
	*User              // 匿名嵌套
	Password *struct{} `json:"pass_word,omitempty"`
}
