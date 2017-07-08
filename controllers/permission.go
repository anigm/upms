package controllers

import (
	"github.com/astaxie/beego"
)

type PermissionController struct {
	beego.Controller
}

func (c *PermissionController) URLMapping() {
	c.Mapping("CreateGroup", c.CreateGroup)
	c.Mapping("DeleteGroup", c.DeleteGroup)
	c.Mapping("CreatePermission", c.CreatePermission)
	c.Mapping("ModifyPermission", c.ModifyPermission)
	c.Mapping("DeletePermission", c.DeletePermission)
	c.Mapping("AddPermission", c.AddPermission)
	c.Mapping("RemovePermission", c.RemovePermission)
	c.Mapping("AddRole", c.AddRole)
	c.Mapping("RemoveRole", c.RemoveRole)
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
// @router /permission/create [put]
func (c *PermissionController) CreatePermission() {
	c.Ctx.WriteString("@router /permission/create [put]")
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
// @router /permission/group/add [post]
func (c *PermissionController) AddPermission() {
	c.Ctx.WriteString("@router /permission/group/add [post]")
}

// * 组内删除角色
// @router /permission/group/remove [post]
func (c *PermissionController) RemovePermission() {
	c.Ctx.WriteString("@router /permission/group/remove [post]")
}

// * 关联权限
// @router /permission/role/add [post]
func (c *PermissionController) AddRole() {
	c.Ctx.WriteString("@router /permission/role/add [post]")
}

// * 取消权限关联
// @router /permission/role/remove [post]
func (c *PermissionController) RemoveRole() {
	c.Ctx.WriteString("@router /permission/role/remove [post]")
}
