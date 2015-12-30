package request

import (
	. "beechat/models/db"
)

type RequestDbModel struct {
}

func (this *RequestDbModel) GetByClientId(clientId int) []Request {
	var requests []Request
	err := DB.Where("clientId = ?", clientId).Find(&requests)
	if err != nil {
		panic(err)
	}
	return requests
}

func (this *RequestDbModel) Add(data Request) {
	_, err := DB.Insert(data)
	if err != nil {
		panic(err)
	}
}
