package models

import (
	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id int
	Content string
}

func init(){
	orm.RegisterModel(new(Log))
}
