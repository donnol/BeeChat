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
	beego.Router("/message/post", &controllers.MessageController{}, "post:Post")
	beego.Router("/message/fetch", &controllers.MessageController{}, "get:Fetch")
	beego.Router("/message/get", &controllers.MessageController{}, "get:Get")
	beego.Router("/request/get", &controllers.RequestController{}, "get:Get")
}
