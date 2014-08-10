package model

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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
	Id          int       `xorm:"int"`
	App         int64     `xorm:"int not null 'app'"`
	Level       string    `xorm:"text not null 'level'"`
	Message     string    `xorm:"text not null 'message'"`
	Module      string    `xorm:"text not null 'module'"`
	RootLogId   string    `xorm:"text not null 'root_log_id'"`
	LogId       string    `xorm:"text not null 'log_id'"`
	ParentLogId string    `xorm:"text 'parent_log_id'"`
	Created     time.Time `xorm:"created"`
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
		fmt.Printf("%s,%v\n", "Fail to unmarshal with logJson", err)
		return
	}
	_, err = x.Insert(logModel)

	if err != nil {
		fmt.Printf("%s,%v\n", "Fail to insert with xorm", err)
		return
	}
	return
}

func WriteLog(logChan *chan string, sysLogLevel string, logLevel string, app int64, message string, module string, logId string, rootLogId string) string {
	if LevelMapping[sysLogLevel] >= LevelMapping[logLevel] {
		parentLogId := logId
		md5String := fmt.Sprintf("%v%v%v", logId, module, string(time.Now().Unix()))
		h := md5.New()
		h.Write([]byte(md5String))
		logId = hex.EncodeToString(h.Sum(nil))

		log := new(Log)
		log.App = app
		log.Level = sysLogLevel
		log.Message = message
		log.Module = module
		log.LogId = logId
		log.RootLogId = rootLogId
		log.ParentLogId = parentLogId
		logJson, _ := json.Marshal(log)
		go WriteLogFile(logId, string(logJson))
		*logChan <- string(logJson)
		return logId
	}
	return ""
}

func Sendlog(logChan *chan string, url string) {
	for {
		log_data := <-(*logChan)
		var log Log
		json.Unmarshal([]byte(log_data), &log)
		request, _ := http.NewRequest("POST", url, bytes.NewReader([]byte(log_data)))
		request.Header.Set("X-Arkors-Application-Client", "127.0.0.1,OAUTH")
		request.Header.Set("Accept", "application/json")
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if resp.StatusCode == http.StatusOK {
			WriteLogStatus(log.LogId, "A")
		} else {
			WriteLogStatus(log.LogId, "X")
		}
	}
}

func WriteLogFile(logId string, log string) {
	filename := string(([]byte(time.Now().Format(time.RFC3339)))[:10]) + ".log"
	f, err := os.OpenFile("logFiles/"+filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.Write([]byte(logId + "|"))
	w.Write([]byte(log + "|"))
	w.Write([]byte("I" + "|"))
	w.Write([]byte("1\n"))
	w.Flush()
}

func WriteLogStatus(logId string, status string) {
	filename := string(([]byte(time.Now().Format(time.RFC3339)))[:10]) + "-receieve.log"
	f, err := os.OpenFile("logFiles/"+filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	w.Write([]byte(logId + "|"))
	w.Write([]byte(status + "\n"))
	w.Flush()
}
