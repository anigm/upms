package controllers

import (
	"ums/models"

	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

var userQuery = &models.UserQuery{}

func (c *LoginController) Login() {
	account := c.GetString("account")
	password := c.GetString("password")
	c.Data["json"] = loginNormal(account, password)
	c.ServeJSON()
}

func loginNormal(account, password string) map[string]interface{} {
	var result = make(map[string]interface{})
	if account == "" && password == "" {
		result["code"] = 4
		result["message"] = "account or password empty"
	} else if user, err := userQuery.GetUserByAccount(account); err == nil {
		if user.Password == "" {
			result["code"] = 1
			result["message"] = "no account:" + account
		} else if user.CheckPassword(password, true) {
			result["code"] = 0
			result["message"] = "login success"
			result["user"] = user
			beego.Info(fmt.Sprintf("%s login success", account))
		} else {
			result["code"] = 2
			result["message"] = "wrong password"
		}
	} else {
		result["code"] = 3
		result["message"] = "error"
		beego.Error(err)
	}
	return result
}
