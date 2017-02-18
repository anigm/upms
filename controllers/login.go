package controllers

import (
	"ums/models"

	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type LoginResult struct {
	Code     int
	Account  string
	UserName string
	Message  string
}

type LoginUser struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

var userQuery = &models.UserQuery{}

func (c *LoginController) Login() {
	account := c.GetString("account")
	password := c.GetString("password")
	if c.Ctx.Input.Header("Content-Type") == "application/json" {
		var user models.User
		models.ParseJson(string(c.Ctx.Input.RequestBody), &user)
		account = user.Account
		password = user.Password
	}
	var result LoginResult
	if user, err := userQuery.GetUserByAccount(account); err == nil {
		if user.Password == "" {
			result.Code = 1
			result.Message = "no account:" + account
		} else if user.CheckPassword(password, true) {
			result.Account = account
			result.UserName = user.UserName
			result.Message = "login success"
			beego.Info(fmt.Sprintf("%s login success", account))
		} else {
			result.Code = 2
			result.Message = "wrong password"
		}
	} else {
		beego.Error(err)
	}
	c.Data["json"] = result
	c.ServeJSON()
}
