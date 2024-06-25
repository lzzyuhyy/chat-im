package models

import (
	"chat-im/global"
	"gorm.io/gorm"
)

// 群表
type GroupChat struct {
	gorm.Model
	OwnerId     uint   `gorm:"type:int(11);comment:当前用户id" json:"owner_id"`
	GroupAvatar string `gorm:"type:varchar(100);comment:群头像" json:"group_avatar"`
	GroupName   string `gorm:"type:varchar(30);comment:群名" json:"group_name"`
}

func (c *GroupChat) GetGroupList() (list []GroupChat, res *gorm.DB) {
	res = global.DB.Table("group_chats").Where("owner_id = ?", c.OwnerId).Find(&list)
	return
}
