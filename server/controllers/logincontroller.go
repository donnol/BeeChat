package controllers

import (
	. "beechat/models/client"
)

type LoginController struct {
	BaseController
	ClientLoginAo ClientLoginAoModel
}

func (this *LoginController) Login() {
	name := this.Ctx.Input.Query("name")
	password := this.Ctx.Input.Query("password")

	result := this.ClientLoginAo.Login(this.Ctx, name, password)

	this.View(result)
}

func (this *LoginController) Logout() {
	this.ClientLoginAo.Logout(this.Ctx)

	result := struct{}{}
	this.View(result)
}

func (this *LoginController) Islogin() {
	result := struct {
		Client
		Islogin bool `json:"islogin"`
	}{}

	client := this.ClientLoginAo.Islogin(this.Ctx)
	if client.ClientId != 0 {
		result.Client = client
		result.Islogin = true
	}
	this.View(result)
}
