package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//Log Level: silly, verbose, info, data, warn, debug, error.
type Log struct {
	Id      int
	Level   string
	Root    string
	Up      string
	Current string
	Message string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Log))
}
