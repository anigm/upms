package routers

import (
	"ums/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/login", &controllers.LoginController{}, "Post:Login"),
	)
	beego.AddNamespace(ns)
}
