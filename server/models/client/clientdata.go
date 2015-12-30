package client

import (
	"time"
)

type Client struct {
	ClientId   int       `json:"clientId" xorm:"autoincr"`
	Name       string    `json:"name"`
	Password   string    `json:"-"`
	CreateTime time.Time `json:"createTime" xorm:"created"`
	ModifyTime time.Time `json:"modifyTime" xorm:"updated"`
}
