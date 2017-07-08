package controllers

import (
	"github.com/astaxie/beego"
)

type PermissionController struct {
	beego.Controller
}

// * 添加分组
// @router /permission/group/create [put]
func (c *PermissionController) CreateGroup() {
	c.Ctx.WriteString("@router /permission/group/create [put]")
}

// * 删除分组
// @router /permission/group/delete [delete]
func (c *PermissionController) DeleteGroup() {
	c.Ctx.WriteString("@router /permission/group/delete [delete]")
}

// * 添加角色
// @router /permission/delete [post]
func (c *PermissionController) CreatePermission() {
	c.Ctx.WriteString("@router /permission/delete [post]")
}

// * 修改角色
// @router /permission/modify [post]
func (c *PermissionController) ModifyPermission() {
	c.Ctx.WriteString("@router /permission/modify [post]")
}

// * 删除角色
// @router /permission/delete [delete]
func (c *PermissionController) DeletePermission() {
	c.Ctx.WriteString("@router /permission/delete [delete]")
}

// * 组内添加角色
// @router /permission/group/add [delete]
func (c *PermissionController) AddPermission() {
	c.Ctx.WriteString("@router /permission/group/add [delete]")
}

// * 组内删除角色
// @router /permission/group/remove [post]
func (c *PermissionController) RemovePermission() {
	c.Ctx.WriteString("@router /permission/group/remove [post]")
}

// * 关联权限
// @router /permission/role/add [post]
func (c *PermissionController) AddRole() {
	c.Ctx.WriteString("@router /permission/permission/add [post]")
}

// * 取消权限关联
// @router /permission/role/remove [post]
func (c *PermissionController) RemoveRole() {
	c.Ctx.WriteString("@router /permission/permission/remove [post]")
}
