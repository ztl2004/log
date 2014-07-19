package model

import (
_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
  "encoding/json"
)

type Log struct {
  Id        int       `xorm:"int pk 'id'"`
  App       int       `xorm:"int not null 'app'"`
  Level     string    `xorm:"text not null 'level'"`
  Action    string    `xorm:"text not null 'action'"`
  FromModel string    `xorm:"text not null 'from_model'"`
  Log       string    `xorm:"text not null 'log'"`
  ParentLog string    `xorm:"text not null 'parentLog'"`
  Created   time.Time `xorm:"created"`
}

var x *xorm.Engine

func init() {
	var err error

	x, err = xorm.NewEngine("mysql", "arkors:arkors@/arkors_log?charset=utf8")

	if err != nil {
		log.Fatalf("Fail to create engine: %v\n", err)
	}

	if err = x.Sync(&Log{}); err != nil {
		log.Fatalf("Fail to sync log database: %v\n", err)
	}

}

func InsertLog(logJson []byte) []byte {
    var logModel Log
    err:=json.Unmarshal(logJson,&logModel)
    if err!=nil {
      log.Fatalf("Fail to unmarshal with logJson",err)
    }
    _,err=x.Insert(logModel)

    if err!=nil {
      log.Fatalf("Fail to insert with xorm",err)
    }

    reModel:=new(ReModel)
    reModel.Id=logModel.Id
    reModel.App=logModel.App
    redata,_:=json.Marshal(reModel)
    return  redata
}
