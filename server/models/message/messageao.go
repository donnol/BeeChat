package message

type MessageAoModel struct {
	MessageDb MessageDbModel
}

func (this *MessageAoModel) Send(sendClientId, receiveClientId int, text string) {
	this.MessageDb.Add(Message{
		SendClientId:    sendClientId,
		ReceiveClientId: receiveClientId,
		Text:            text,
		Type:            1, //暂定为未读信息
	})
}

func (this *MessageAoModel) Recv(clientId int) []Message {
	result := this.MessageDb.GetUnreadByReceiveClientId(clientId)
	return result
}
