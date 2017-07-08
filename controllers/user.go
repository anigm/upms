package controllers

import (
	"strconv"
	"upms/models"

	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type LoginResult struct {
	Code    int
	Name    string
	Account string
	Success bool
}

// * 注册
// * 登陆
// * 删除
// * 查询
// * 修改信息
// * 修改密码

func (c *UserController) Login() {
	account := c.GetString("account")
	password := c.GetString("password")
	beego.Debug(account, password)
	if user, err := models.GetUserByAccount(account); err == nil {
		if user.Password == password {
			var result LoginResult
			result.Account = account
			result.Name = user.Name
			result.Success = true
			c.Data["json"] = result
			c.ServeJSON()
			return
		}
	} else {
		beego.Error(err)
	}
	var result LoginResult
	result.Success = false
	c.Data["json"] = result
	c.ServeJSON()
}
func (c *UserController) Get() {
	users, _ := models.GetUsers()
	c.Ctx.Output.Header("Content-type", "application/json;charset=utf-8")
	c.Ctx.WriteString(models.DBModelToJson(users))
}

func (c *UserController) GetUserInfo() {
	idstr := c.Ctx.Input.Param(":id")
	if id, err := strconv.ParseInt(idstr, 0, 64); err == nil {
		user, _ := models.GetUserInfo(id)
		c.Ctx.Output.Header("Content-type", "application/json;charset=utf-8")
		c.Ctx.WriteString(models.DBModelToJson(user))
	}
}

func (c *UserController) GetUserInfoByQQ() {
	qq := c.Ctx.Input.Param(":qq")
	user, _ := models.GetUserInfoByQQ(qq)
	c.Ctx.Output.Header("Content-type", "application/json;charset=utf-8")
	c.Ctx.WriteString(models.DBModelToJson(user))
}

type CreateUserResult struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}

func (c *UserController) CreateUser() {
	var name, qq string
	c.Ctx.Input.Bind(&name, "name")
	c.Ctx.Input.Bind(&qq, "qq")
	beego.Debug(name)
	beego.Debug(qq)
	if name == "" || qq == "" {
		c.Data["json"] = &JsonResult{Status: -1, Info: "参数错误"}
		c.ServeJSON()
	}
	var user models.User
	user.Name = name
	user.QQ = qq
	if err := models.SaveUser(&user); err == nil {
		c.Data["json"] = &JsonResult{Status: 0, Info: "success"}
		c.ServeJSON()
	} else {
		c.Data["json"] = &JsonResult{Status: -1, Info: err.Error()}
		c.ServeJSON()
	}
}

func (c *UserController) DeleteUser() {
	ids := strings.Split(string(c.Ctx.Input.RequestBody), ",")
	beego.Debug(ids)
	if err := models.DeleteUsers(ids); err == nil {
		c.Data["json"] = &JsonResult{Status: 0, Info: "success"}
		c.ServeJSON()
	} else {
		c.Data["json"] = &JsonResult{Status: -1, Info: err.Error()}
		c.ServeJSON()
	}
}
