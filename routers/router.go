package routers

import (
	"3arrows_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/dologin", &controllers.UserController{}, "post:DoLogin")
}
