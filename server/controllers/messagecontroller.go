package controllers

import (
	. "beechat/models/archive"
	. "beechat/models/chatroom"
	. "beechat/models/client"
	. "beechat/models/message"
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

	Join(client.Name)
	Publish <- NewEvent(EVENT_MESSAGE, client.Name, text)

	result := struct{}{}
	this.View(result)
}

func (this *MessageController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	if err != nil {
		return
	}

	this.ClientLoginAo.CheckMustLogin(this.Ctx)

	events := GetEvents(int(lastReceived))
	if len(events) > 0 {
		this.Data["json"] = events
		this.ServeJson()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	WaitingList.PushBack(ch)
	<-ch

	this.Data["json"] = GetEvents(int(lastReceived))
	this.ServeJson()
}
