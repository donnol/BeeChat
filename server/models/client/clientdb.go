package client

import (
	. "beechat/models/db"
	"fmt"
)

type ClientDbModel struct {
}

func (this *ClientDbModel) Get(id int) Client {
	var clients []Client
	err := DB.Where("clientId = ?", id).Find(&clients)
	fmt.Println(id)
	if err != nil {
		panic(err)
	}
	if len(clients) == 0 {
		panic("该用户不存在！")
	}
	fmt.Println(clients[0])
	return clients[0]
}

func (this *ClientDbModel) GetByNameAndPassword(name string, password string) []Client {
	var clients []Client
	err := DB.Where("name = ? and password = ?", name, password).Find(&clients)
	if err != nil {
		panic(err)
	}
	return clients
}
