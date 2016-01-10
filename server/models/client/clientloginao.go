package client

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/session"

	. "beechat/models/chatroom"
)

var globalSessions *session.Manager

type ClientLoginAoModel struct {
	ClientAo ClientAoModel
}

func (this *ClientLoginAoModel) sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func (this *ClientLoginAoModel) Login(ctx *context.Context, name string, password string) Client {
	password = this.sha1(password)
	data := this.ClientAo.GetByNameAndPassword(name, password)
	if len(data) == 0 {
		panic("login failed, don't exist client!")
	} else {
		ctx.Output.Session("name", data[0].ClientId)
		Join(data[0].Name)
		return data[0]
	}
}

func (this *ClientLoginAoModel) Logout(ctx *context.Context) {
	clientId, ok := ctx.Input.Session("name").(int)
	if ok == true && clientId != 0 {
		client := this.ClientAo.Get(clientId)
		Leave(client.Name)
		ctx.Output.Session("name", 0)
		fmt.Println("logout!")
	}
}

func (this *ClientLoginAoModel) Islogin(ctx *context.Context) Client {
	var result Client
	clientId, ok := ctx.Input.Session("name").(int)
	if ok == true && clientId > 10000 {
		result = this.ClientAo.Get(clientId)
	} else {
		result = Client{}
	}
	return result
}

func (this *ClientLoginAoModel) CheckMustLogin(ctx *context.Context) Client {
	client := this.Islogin(ctx)
	if client.ClientId == 0 {
		panic("用户未登陆！")
	}
	return client
}
