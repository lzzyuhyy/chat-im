package models

import (
	"chat-im/global"
	"fmt"
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	var err error
	host := beego.AppConfig.String("mysql_host")
	port, _ := beego.AppConfig.Int("mysql_port")
	user := beego.AppConfig.String("mysql_user")
	pass := beego.AppConfig.String("mysql_pass")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/impro?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	global.DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("数据库链接失败", err)
		return
	}

	global.DB.AutoMigrate(new(User), new(UserRelationship), new(MsgHistories), new(GroupChat))
}
