package request

import (
	"time"
)

type Request struct {
	RequestId   int       `json:"requestId" xorm:"autoincr"`
	ClientId    int       `json:"clientId"`
	RequestTime time.Time `json:"requestTime"`
	CreateTime  time.Time `json:"createTime" xorm:"created"`
	ModifyTime  time.Time `json:"modifyTime" xorm:"updated"`
}
