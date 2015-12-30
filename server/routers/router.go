package routers

import (
	"beechat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get,post:Get")
	beego.Router("/client/get", &controllers.ClientController{}, "get:Get")
	beego.Router("/login/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/login/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/login/islogin", &controllers.LoginController{}, "get:Islogin")
	beego.Router("/message/send", &controllers.MessageController{}, "post:Send")
	beego.Router("/message/recv", &controllers.MessageController{}, "get:Recv")
	beego.Router("/request/get", &controllers.RequestController{}, "get:Get")
}
