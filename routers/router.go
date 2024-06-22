package routers

import (
	"chat-im/controllers"
	"chat-im/routers/userRouter"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/v1/chan", &controllers.ChanController{}, "GET:Chan")

	v1 := beego.NewNamespace("/api/v1",
		userRouter.UserRouter(),
	)

	beego.AddNamespace(v1)
}
