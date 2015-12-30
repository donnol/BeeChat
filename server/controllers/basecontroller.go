package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) View(result interface{}) {
	this.Data["json"] = &result
	this.ServeJson()
}
