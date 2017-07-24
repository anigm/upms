package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "CreateGroup",
			Router: `/permission/group/create`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "DeleteGroup",
			Router: `/permission/group/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "CreatePermission",
			Router: `/permission/create`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "ModifyPermission",
			Router: `/permission/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "DeletePermission",
			Router: `/permission/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "AddPermission",
			Router: `/permission/group/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "RemovePermission",
			Router: `/permission/group/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "AddRole",
			Router: `/permission/role/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:PermissionController"] = append(beego.GlobalControllerRouter["upms/controllers:PermissionController"],
		beego.ControllerComments{
			Method: "RemoveRole",
			Router: `/permission/role/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "CreateGroup",
			Router: `/role/group/create`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "DeleteGroup",
			Router: `/role/group/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "CreateRole",
			Router: `/role/create`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "ModifyRole",
			Router: `/role/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "DeleteRole",
			Router: `/role/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "AddRole",
			Router: `/role/group/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RemoveRole",
			Router: `/role/group/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "AddPermission",
			Router: `/role/permission/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:RoleController"] = append(beego.GlobalControllerRouter["upms/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RemovePermission",
			Router: `/role/permission/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "DeleteUser",
			Router: `/user/delete`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "User",
			Router: `/user`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "User",
			Router: `/user/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "ModifyUser",
			Router: `/user/modify`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "Passwd",
			Router: `/user/passwd`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "AddRole",
			Router: `/user/role/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["upms/controllers:UserController"] = append(beego.GlobalControllerRouter["upms/controllers:UserController"],
		beego.ControllerComments{
			Method: "RemoveRole",
			Router: `/user/role/remove`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
