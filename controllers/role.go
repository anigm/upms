package controllers

import (
	"github.com/astaxie/beego"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) URLMapping() {
	c.Mapping("CreateGroup", c.CreateGroup)
	c.Mapping("DeleteGroup", c.DeleteGroup)
	c.Mapping("CreateRole", c.CreateRole)
	c.Mapping("ModifyRole", c.ModifyRole)
	c.Mapping("DeleteRole", c.DeleteRole)
	c.Mapping("AddRole", c.AddRole)
	c.Mapping("RemoveRole", c.RemoveRole)
	c.Mapping("AddPermission", c.AddPermission)
	c.Mapping("RemovePermission", c.RemovePermission)
}

// * 添加分组
// @router /role/group/create [put]
func (c *RoleController) CreateGroup() {
	c.Ctx.WriteString("@router /role/group/create [put]")
}

// * 删除分组
// @router /role/group/delete [delete]
func (c *RoleController) DeleteGroup() {
	c.Ctx.WriteString("@router /role/group/delete [delete]")
}

// * 添加角色
// @router /role/create [put]
func (c *RoleController) CreateRole() {
	c.Ctx.WriteString("@router /role/create [put]")
}

// * 修改角色
// @router /role/modify [post]
func (c *RoleController) ModifyRole() {
	c.Ctx.WriteString("@router /role/modify [post]")
}

// * 删除角色
// @router /role/delete [delete]
func (c *RoleController) DeleteRole() {
	c.Ctx.WriteString("@router /role/delete [delete]")
}

// * 组内添加角色
// @router /role/group/add [post]
func (c *RoleController) AddRole() {
	c.Ctx.WriteString("@router /role/group/add [post]")
}

// * 组内删除角色
// @router /role/group/remove [post]
func (c *RoleController) RemoveRole() {
	c.Ctx.WriteString("@router /role/group/remove [post]")
}

// * 关联权限
// @router /role/permission/add [post]
func (c *RoleController) AddPermission() {
	c.Ctx.WriteString("@router /role/permission/add [post]")
}

// * 取消权限关联
// @router /role/permission/remove [post]
func (c *RoleController) RemovePermission() {
	c.Ctx.WriteString("@router /role/permission/remove [post]")
}
