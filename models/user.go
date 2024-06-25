package models

import (
	"chat-im/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account  string `gorm:"type:varchar(20);unique;comment:账号" json:"account,omitempty"`
	Nickname string `gorm:"type:varchar(50);comment:昵称" json:"nickname,omitempty"`
	Password string `gorm:"type:char(32);comment:密码" json:"password,omitempty"`
	Salt     string `gorm:"type:varchar(10);comment:盐" json:"salt,omitempty"`
	Mobile   string `gorm:"type:char(11);unique;comment:手机号" json:"mobile,omitempty"`
	Email    string `gorm:"type:varchar(100);comment:邮箱" json:"email,omitempty"`
	Avatar   string `gorm:"type:varchar(150);comment:头像" json:"avatar,omitempty"`
	Status   int8   `gorm:"type:tinyint(1);comment:账号状态" json:"status,omitempty"`
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

func (u *User) FindUserByEmail() {
	global.DB.Table("users").Where("email = ?", u.Email).Limit(1).Find(&u)
}
