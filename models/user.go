package models

import (
	"chat-im/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `gorm:"type:varchar(20);unique;comment:账号"`
	Nickname string `gorm:"type:varchar(50);comment:昵称"`
	Password string `gorm:"type:char(32);comment:密码"`
	Salt     string `gorm:"type:varchar(10);comment:盐"`
	Mobile   string `gorm:"type:char(11);unique;comment:手机号"`
	Email    string `gorm:"type:varchar(100);comment:邮箱"`
	Avatar   string `gorm:"type:varchar(150);comment:头像"`
	Status   int8   `gorm:"type:tinyint(1);comment:账号状态"`
}

func (u *User) FindUser() (User, *gorm.DB) {
	var userInfo User
	res := global.DB.Table("users").Where("account = ?", u.Account).Limit(1).Find(&userInfo)
	return userInfo, res
}

func (u *User) CreateUser() *gorm.DB {
	return global.DB.Table("users").Create(&u)
}

// 获取用户信息
func (u *User) GetUserInfoById() *gorm.DB {
	return global.DB.Table("users").Limit(1).Find(&u)
}
