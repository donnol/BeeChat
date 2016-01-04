package controllers

import (
	"github.com/astaxie/beego"
)

const (
	host = "www.test2.hongbeibang.com"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	inHost := this.Ctx.Input.Host()
	if inHost != host {
		panic("非法请求！")
	}
	port := this.Ctx.Input.Port()
	if port != 9999 {
		panic("非法端口！")
	}
}

func (this *BaseController) View(result interface{}) {
	this.Data["json"] = &result
	this.ServeJson()
}
