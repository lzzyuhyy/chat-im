package models

import (
	"chat-im/global"
	"gorm.io/gorm"
)

type MsgHistories struct {
	gorm.Model
	OwnerId uint   `gorm:"type:int(11);comment:主端id" json:"owner_id"`
	DistId  uint   `gorm:"type:int(11);comment:对端id" json:"dist_id"`
	Content string `gorm:"type:varchar(100);comment:消息内容" json:"content"`
	IsRead  uint8  `gorm:"type:tinyint(1);comment:读取状态" json:"is_read"`
	IsSend  uint8  `gorm:"type:tinyint(1);comment:发送状态" json:"is_send"`
	Cmd     uint8  `gorm:"type:tinyint(1);comment:消息类型 1私聊 2群聊" json:"cmd"`
	Status  uint8  `gorm:"type:tinyint(1);comment:消息状态" json:"status"`
}

// 添加聊天记录
func (mh *MsgHistories) AddMessage() error {
	return global.DB.Create(&mh).Error
}
