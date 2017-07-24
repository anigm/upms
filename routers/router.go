package routers

import (
	"upms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// beego.NSRouter("/user", &controllers.UserController{}),
		// beego.NSRouter("/user/:id", &controllers.UserController{}, "GET:GetUserInfo"),
		// beego.NSRouter("/login", &controllers.UserController{}, "GET:Login"),
		// beego.NSRouter("/user/qq/:qq", &controllers.UserController{}, "GET:GetUserInfoByQQ"),
		// beego.NSRouter("/user/create", &controllers.UserController{}, "Post:CreateUser"),
		// beego.NSRouter("/user/delete", &controllers.UserController{}, "Post:DeleteUser"),
		beego.NSInclude(&controllers.RoleController{}),
		beego.NSInclude(&controllers.PermissionController{}),
		beego.NSInclude(&controllers.UserController{}),
	)
	beego.AddNamespace(ns)
}
