package controllers

import (
	_ "upms/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

const (
	ResultStatus_Success = 0
	ResultStatus_Error   = -1
)

type JsonResult struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data,omitempty"`
}
