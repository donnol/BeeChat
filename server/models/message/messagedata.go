package message

import (
	"time"
)

type Message struct {
	MessageId       int       `json:"messageId" xorm:"autoincr"`
	Text            string    `json:"text"`
	SendClientId    int       `json:"sendClientId"`
	ReceiveClientId int       `json:"receiveClientId"`
	Type            int       `json:"type"`
	CreateTime      time.Time `json:"createTime" xorm:"created"`
	ModifyTime      time.Time `json:"modifyTime" xorm:"updated"`
}
