package userRouter

import (
	"chat-im/controllers"
	"github.com/astaxie/beego"
)

func UserRouter() beego.LinkNamespace {
	return beego.NSNamespace("/user",
		beego.NSRouter("/register", &controllers.UserController{}, "POST:Register"),
		beego.NSRouter("/login", &controllers.UserController{}, "POST:Login"),
		beego.NSRouter("/friend/list", &controllers.UserController{}, "Get:FriendList"),
	)
}
