package models

import (
	"chat-im/global"
	"gorm.io/gorm"
)

type UserRelationship struct {
	gorm.Model
	OwnerId uint  `gorm:"type:int(11);comment:主端id" json:"owner_id"`
	DistId  uint  `gorm:"type:int(11);comment:对端id" json:"dist_id"`
	Status  uint8 `gorm:"type:tinyint(1);comment:关系状态0-待确认 1-好友 2拉黑" json:"status"`
	IsGroup uint8 `gorm:"type:tinyint(1);comment:是否群聊0-否 1-是" json:"is_group"`
}

// 好友列表结构
type FriendList struct {
	UserRelationship
	UserInfo User `json:"user_info"`
}

func (ur *UserRelationship) GetUserDistId() (list []UserRelationship, res *gorm.DB) {
	res = global.DB.Table("user_relationships").Where("owner_id = ? and status = 1", ur.OwnerId).Find(&list)
	return
}

// 通过群id查在群里的所有用户
func (ur *UserRelationship) GetUserByGroupId() (list []UserRelationship, res *gorm.DB) {
	res = global.DB.Table("user_relationships").Where("dist_id= ? and is_group = 1", ur.DistId).Find(&list)
	return
}
