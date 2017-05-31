package routers

import (
	"upms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/user", &controllers.UserController{}),
		beego.NSRouter("/user/:id", &controllers.UserController{}, "GET:GetUserInfo"),
		beego.NSRouter("/login", &controllers.UserController{}, "GET:Login"),
		beego.NSRouter("/user/qq/:qq", &controllers.UserController{}, "GET:GetUserInfoByQQ"),
		beego.NSRouter("/user/create", &controllers.UserController{}, "Post:CreateUser"),
		beego.NSRouter("/user/delete", &controllers.UserController{}, "Post:DeleteUser"),
		beego.NSRouter("/platforms", &controllers.PlatformController{}, "GET:Platforms"),
		beego.NSRouter("/roles", &controllers.RoleController{}, "Post:Roles"),
		beego.NSRouter("/rolegroups", &controllers.RoleController{}, "Get:RoleGroups"),
	)
	beego.AddNamespace(ns)
}
