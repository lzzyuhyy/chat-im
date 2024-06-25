package main

import (
	"chat-im/models"
	_ "chat-im/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*", "token"},
	}))

	models.InitMySQL()
	beego.Run()
}
