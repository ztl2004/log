package model

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type LEVEL int32

const (
	OFF LEVEL = iota
	ERROR
	WARN
	INFO
	DEBUG
	ALL
)

var LevelMapping = map[string]LEVEL{
	"OFF":   OFF,
	"ERROR": ERROR,
	"WARN":  WARN,
	"INFO":  INFO,
	"DEBUG": DEBUG,
	"ALL":   ALL,
}

type Log struct {
	Id        int       `xorm:"int"`
	App       int64     `xorm:"int not null 'app'"`
	Level     string    `xorm:"text not null 'level'"`
	Action    string    `xorm:"text not null 'action'"`
	Module    string    `xorm:"text not null 'module'"`
	Log       string    `xorm:"text not null 'log'"`
	ParentLog string    `xorm:"text 'parentLog'"`
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

func InsertLog(logJson []byte) {
	var logModel Log
	err := json.Unmarshal(logJson, &logModel)
	if err != nil {
		log.Fatalf("Fail to unmarshal with logJson", err)
	}
	_, err = x.Insert(logModel)

	if err != nil {
		log.Fatalf("Fail to insert with xorm", err)
	}

	return
}

func WriteLog(logChan *chan string, sysLogLevel string, logLevel string, app int64, action string, module string, logId string) string {
	if LevelMapping[sysLogLevel] >= LevelMapping[logLevel] {
		parentLogId := logId
		md5String := fmt.Sprintf("%v%v%v", logId, module, string(time.Now().Unix()))
		h := md5.New()
		h.Write([]byte(md5String))
		logId = hex.EncodeToString(h.Sum(nil))

		log := new(Log)
		log.App = app
		log.Level = sysLogLevel
		log.Action = action
		log.Module = module
		log.Log = logId
		log.ParentLog = parentLogId
		logJson, _ := json.Marshal(log)
		*logChan <- string(logJson)
		return logId
	}
	return ""
}

func Sendlog(logChan *chan string) {
	for {
		log_data := <-(*logChan)
		var log Log
		json.Unmarshal([]byte(log_data), &log)
		request, _ := http.NewRequest("POST", "http://log.arkors.com/v1/log", bytes.NewReader([]byte(log_data)))
		request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,OAUTH")
		request.Header.Set("Accept", "application/json")
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if resp.StatusCode == http.StatusOK {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(data))
		}
	}
}
