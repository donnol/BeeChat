package controllers

import (
	. "beechat/models/client"
	. "beechat/models/message"
	"fmt"
	"strconv"
)

type MessageController struct {
	BaseController
	MessageAo     MessageAoModel
	ClientAo      ClientAoModel
	ClientLoginAo ClientLoginAoModel
}

func (this *MessageController) Send() {
	receiveClientId := this.Ctx.Input.Query("receiveClientId")
	text := this.Ctx.Input.Query("text")

	client := this.ClientLoginAo.CheckMustLogin(this.Ctx)

	receClientId, err := strconv.Atoi(receiveClientId)
	if err != nil {
		panic(err)
	}
	this.MessageAo.Send(client.ClientId, receClientId, text)

	result := struct{}{}
	this.View(result)
}

func (this *MessageController) Recv() {
	client := this.ClientLoginAo.CheckMustLogin(this.Ctx)
	clientId := client.ClientId

	result := this.MessageAo.Recv(clientId)
	fmt.Println(result)
	this.View(result)
}
