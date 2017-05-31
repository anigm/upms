package controllers

import (
	"strings"
	"upms/models"

	"github.com/astaxie/beego"
)

type RoleController struct {
	beego.Controller
}

type RolesData struct {
	ID       string
	Role     string
	Platform string
}

func (c *RoleController) Roles() {
	psstr := string(c.Ctx.Input.RequestBody)
	var platforms []string
	var datas []RolesData
	platforms = strings.Split(psstr, ",")
	if db, err := models.DB(); err == nil {
		if rows, err := db.Raw("SELECT id,name,platform FROM roles where platform in ('" + strings.Join(platforms, "','") + "')").Rows(); err == nil {
			defer rows.Close()
			for rows.Next() {
				var id string
				var platform string
				var role string
				rows.Scan(&id, &role, &platform)
				datas = append(datas, RolesData{ID: id, Role: role, Platform: platform})
			}
		} else {
			c.Data["json"] = &JsonResult{Status: ResultStatus_Error, Info: err.Error()}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = &JsonResult{Status: ResultStatus_Error, Info: err.Error()}
		c.ServeJSON()
	}
	c.Data["json"] = &JsonResult{Status: ResultStatus_Success, Info: "success", Data: datas}
	c.ServeJSON()
}

func (c *RoleController) RoleGroups() {
	var roleGroups []models.RoleGroup
	if db, err := models.DB(); err == nil {
		if rows, err := db.Raw("SELECT id,name,parent_id FROM role_groups order by id").Rows(); err == nil {
			defer rows.Close()
			for rows.Next() {
				var r models.RoleGroup
				rows.Scan(&r.ID, &r.Name, &r.ParentID)
				db.Model(r).Select("id,name,platform").Association("Roles").Find(&r.Roles)
				roleGroups = append(roleGroups, r)
			}
		} else {
			c.Data["json"] = &JsonResult{Status: ResultStatus_Error, Info: err.Error()}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = &JsonResult{Status: ResultStatus_Error, Info: err.Error()}
		c.ServeJSON()
	}
	c.Data["json"] = &JsonResult{Status: ResultStatus_Success, Info: "success", Data: roleGroups}
	c.ServeJSON()
}
