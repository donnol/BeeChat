package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

const (
	host = "www.jdscript.com"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	inHost := this.Ctx.Input.Host()
	fmt.Println(inHost)
	if inHost != host {
		panic("非法请求！")
	}
	port := this.Ctx.Input.Port()
	fmt.Println(port)
	if port != 9999 {
		panic("非法端口！")
	}
}

func (this *BaseController) View(result interface{}) {
	this.Data["json"] = &result
	this.ServeJson()
}
