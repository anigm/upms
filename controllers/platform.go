package controllers

import "github.com/astaxie/beego"
import "upms/models"

type PlatformController struct {
	beego.Controller
}

func (c *PlatformController) Platforms() {
	var platforms []string
	if db, err := models.DB(); err == nil {
		if rows, err := db.Raw("SELECT distinct(platform) FROM roles").Rows(); err == nil {
			defer rows.Close()
			for rows.Next() {
				var platform string
				rows.Scan(&platform)
				platforms = append(platforms, platform)
			}
		} else {
			c.Data["json"] = &JsonResult{Status: ResultStatus_Success, Info: "success"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = &JsonResult{Status: ResultStatus_Error, Info: err.Error()}
		c.ServeJSON()
	}
	c.Data["json"] = &JsonResult{Status: ResultStatus_Success, Info: "success", Data: platforms}
	c.ServeJSON()
}
