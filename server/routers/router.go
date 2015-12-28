package routers

import (
	"BeeChat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get,post:Get")
}
