package controllers

import (
	"3arrows_web/models"
	"fmt"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	UserService *models.UserService
}

//初始化?
func init() {
	fmt.Print("\r\n初始化BaseController\r\n")
}

//检查参数
func (this *BaseController) CheckParams(params []string) bool {
	for _, v := range params {
		if len(this.GetString(v)) == 0 {
			this.Ctx.WriteString(fmt.Sprintf("missing params of %s", v))
			return false
		}
	}
	return true
}

//返回json
func (this *BaseController) Json(data interface{}, msg string) bool {
	if data == nil && len(msg) == 0 {
		this.Ctx.WriteString("")
		return true
	}
	r := map[string]interface{}{"data": &data, "msg": msg}

	this.Data["json"] = &r
	this.ServeJSON()
	return true
}
