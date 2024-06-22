package main

import (
	"chat-im/models"
	_ "chat-im/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.InitMySQL()
	beego.Run()
}
