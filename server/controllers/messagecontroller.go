package controllers

import (
	. "beechat/models/archive"
	. "beechat/models/chatroom"
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

func (this *MessageController) Get() {
	client := this.ClientLoginAo.CheckMustLogin(this.Ctx)

	result := this.MessageAo.GetByClientId(client.ClientId)
	this.View(result)
}

func (this *MessageController) Post() {
	receiveClientId := this.Ctx.Input.Query("receiveClientId")
	text := this.Ctx.Input.Query("text")

	client := this.ClientLoginAo.CheckMustLogin(this.Ctx)

	receClientId, err := strconv.Atoi(receiveClientId)
	if err != nil {
		receClientId = 0
	}
	this.MessageAo.Send(client.ClientId, receClientId, text)

	result := struct{}{}
	this.View(result)
}

func (this *MessageController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	if err != nil {
		return
	}

	client := this.ClientLoginAo.CheckMustLogin(this.Ctx)
	fmt.Println("login client id: " + strconv.Itoa(client.ClientId) + "login client name: " + client.Name)

	events := GetEvents(int(lastReceived), client.Name)
	if len(events) > 0 {
		fmt.Println("go")
		this.View(events)
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	WaitingList.PushBack(ch)
	<-ch

	this.View(GetEvents(int(lastReceived), client.Name))
}
