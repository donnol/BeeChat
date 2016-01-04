package message

import (
	. "beechat/models/request"
	"time"
)

const (
	timePattern = "2006-01-02 15:04:05"
)

type MessageAoModel struct {
	MessageDb MessageDbModel
	RequestAo RequestAoModel
}

func (this *MessageAoModel) GetByClientId(clientId int) []Message {
	result := this.MessageDb.GetByClientId(clientId)

	return result
}

func (this *MessageAoModel) Send(sendClientId, receiveClientId int, text string) {
	if sendClientId == receiveClientId {
		return
	}

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
