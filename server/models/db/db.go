package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var DB *xorm.Engine

func init() {
	var err error
	DB, err = xorm.NewEngine("mysql", "root:1@/chat?charset=utf8")
	if err != nil {
		panic(err)
	}

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "t_")
	DB.SetTableMapper(tbMapper)

	DB.SetColumnMapper(core.SameMapper{})

	DB.SetMaxOpenConns(10)
}
