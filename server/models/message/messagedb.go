package message

import (
	. "beechat/models/db"
)

type MessageDbModel struct {
}

func (this *MessageDbModel) Get(id int) Message {
	var messages []Message
	err := DB.Where("messageId = ?", id).Find(&messages)
	if err != nil {
		panic(err)
	}
	if len(messages) == 0 {
		panic("不存在该数据！")
	}
	return messages[0]
}

func (this *MessageDbModel) GetUnreadByReceiveClientId(receiveClientId int) []Message {
	var messages []Message
	err := DB.Where("(receiveClientId = ? or receiveClientId = 0) and type = 1", receiveClientId).Find(&messages)
	if err != nil {
		panic(err)
	}
	return messages
}

func (this *MessageDbModel) GetByClientId(clientId int) []Message {
	var messages []Message
	err := DB.Where("sendClientId = ? or receiveClientId = ? or receiveClientId = 0", clientId, clientId).OrderBy("createTime asc").Find(&messages)
	if err != nil {
		panic(err)
	}
	return messages
}

func (this *MessageDbModel) GetBySendClientId(sendClientId int) []Message {
	var messages []Message
	err := DB.Where("sendClientId = ?", sendClientId).Find(&messages)
	if err != nil {
		panic(err)
	}
	return messages
}

func (this *MessageDbModel) Add(data Message) {
	_, err := DB.Insert(data)
	if err != nil {
		panic(err)
	}
}

func (this *MessageDbModel) Mod(id int, data Message) {
	_, err := DB.Where("messageId = ?", id).Update(&data)
	if err != nil {
		panic(err)
	}
}
