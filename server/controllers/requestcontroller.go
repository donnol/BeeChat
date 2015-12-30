package controllers

import (
	. "beechat/models/client"
	. "beechat/models/message"
	. "beechat/models/request"
)

type RequestController struct {
	BaseController
	ClientAo      ClientAoModel
	ClientLoginAo ClientLoginAoModel
	RequestAo     RequestAoModel
	MessageAo     MessageAoModel
}

func (this *RequestController) Get() {
	client := this.ClientLoginAo.CheckMustLogin(this.Ctx)

	result := this.RequestAo.GetByClientId(client.ClientId)
	this.View(result)
}
