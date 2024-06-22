package models

import (
	"chat-im/global"
	"gorm.io/gorm"
)

type UserRelationship struct {
	gorm.Model
	OwnerId uint  `gorm:"type:int(11);comment:主端id"`
	DistId  uint  `gorm:"type:int(11);comment:对端id"`
	Status  uint8 `gorm:"type:tinyint(1);comment:关系状态0-待确认 1-好友 2拉黑"`
	IsGroup uint8 `gorm:"type:tinyint(1);comment:是否群聊0-否 1-是"`
}

// 好友列表结构
type FriendList struct {
	UserRelationship
	UserInfo User
}

func (ur *UserRelationship) GetUserDistId() (list []UserRelationship, res *gorm.DB) {
	res = global.DB.Table("user_relationships").Where("owner_id = ? and status = 1", ur.OwnerId).Find(&list)
	return
}
