package controllers

import (
	"3arrows_web/models"
	"fmt"
	"strconv"
)

type UserController struct {
	BaseController
}

//登录页面
func (this *UserController) Login() {
	this.TplName = "user/login.tpl"
}

//登录ajax请求
func (this *UserController) DoLogin() {
	if !this.CheckParams([]string{"email", "password"}) {
		return
	}
	this.UserService = models.UserServiceInstance()
	email := this.GetString("email")
	password := this.GetString("password")

	userObj, err := this.UserService.CheckPassword(email, password)
	if err != nil {
		this.Json(nil, fmt.Sprintf("%s", err))
		return
	}
	//写入cookie
	this.SetSecureCookie(userObj.CookieToken, "uid", strconv.Itoa(userObj.Uid))
	this.SetSecureCookie(userObj.CookieToken, "t", userObj.LoginToken)

	this.Json(map[string]string{"jump": "/user/index"}, "ok")
}
