package main

import (
	_ "beechat/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.SessionOn = true
	beego.SetStaticPath("/chat", "../static/src/pages")
	beego.SessionProvider = "redis"
	beego.SessionSavePath = "127.0.0.1:6379,100,13420693396"
	beego.Run()
}
