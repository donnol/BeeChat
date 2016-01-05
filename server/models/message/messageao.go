package message

import (
	. "beechat/models/archive"
	. "beechat/models/chatroom"
	. "beechat/models/client"
	. "beechat/models/request"

	"time"
)

const (
	timePattern = "2006-01-02 15:04:05"
)

type MessageAoModel struct {
	MessageDb MessageDbModel
	RequestAo RequestAoModel
	ClientAo  ClientAoModel
}

func (this *MessageAoModel) GetByClientId(clientId int) []Message {
	result := this.MessageDb.GetByClientId(clientId)

	return result
}

func (this *MessageAoModel) Send(sendClientId, receiveClientId int, text string) {
	if sendClientId == receiveClientId {
		return
	}
	receiveClient := Client{}
	if receiveClientId != 0 {
		receiveClient = this.ClientAo.Get(receiveClientId)
	} else {
		receiveClient = Client{
			Name: "all",
		}
	}
	sendClient := this.ClientAo.Get(sendClientId)

	text = sendClient.Name + ": " + text
	Publish <- NewEvent(EVENT_MESSAGE, sendClient.Name, receiveClient.Name, text)

	this.MessageDb.Add(Message{
		SendClientId:    sendClientId,
		ReceiveClientId: receiveClientId,
		Text:            text,
		Type:            1, //暂定为未读信息
	})
}

func (this *MessageAoModel) Recv(clientId int) []Message {
	result := this.MessageDb.GetUnreadByReceiveClientId(clientId)
	for _, value := range result {
		this.MessageDb.Mod(value.MessageId, Message{
			Type: 2,
		})
	}

	in30Minute := time.Now().Unix() - 30*60
	in30MinuteStr := time.Unix(in30Minute, 0).Format(timePattern)
	beginTime, err := time.Parse(timePattern, in30MinuteStr)
	if err != nil {
		panic(err)
	}
	request := this.RequestAo.GetByClientIdAndTime(clientId, beginTime)
	if len(request) == 0 {
		this.RequestAo.Add(clientId)
	}

	return result
}
