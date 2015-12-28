package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	a := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		"jd",
		18,
	}
	this.Data["json"] = &a
	this.ServeJson()
}
