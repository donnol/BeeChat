package controllers

import (
	. "beechat/models/client"
	"strconv"
)

type ClientController struct {
	BaseController
	ClientAo ClientAoModel
}

func (this *ClientController) Get() {
	id := this.Ctx.Input.Query("clientId")
	clientId, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	client := this.ClientAo.Get(clientId)
	this.View(client)
}
