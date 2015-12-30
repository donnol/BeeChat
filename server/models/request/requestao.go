package request

import (
	"time"
)

type RequestAoModel struct {
	RequestDb RequestDbModel
}

func (this *RequestAoModel) GetByClientId(clientId int) []Request {
	return this.RequestDb.GetByClientId(clientId)
}

func (this *RequestAoModel) Add(clientId int) {
	this.RequestDb.Add(Request{
		ClientId:    clientId,
		RequestTime: time.Now(),
	})
}
