package controllers

import (
	"github.com/astaxie/beego"
)

type LogController struct {
	beego.Controller
}

// POST /v1/logs
// 创建 Log 日志
// Example Request:
//    POST /v1/logs HTTP/1.1
//    Host: log.arkors.com
//    X-Arkors-Application-Id: demo
//    X-Arkors-Application-Token: 13b0a8dbddd7c98499a12976ab023ece,1389085779854
//    Accept: application/json
//    {
//      "level": "info",
//      "root": "pnktnjyb996sj4p156gjtp4im",
//      "up": "pnktnjyb996sj4p156gjtp4im",
//      "message": "Create user named genedna"
//    }
// Example Response:
//    HTTP/1.1 201 OK
//    Vary: Accept
//    Content-Type: application/json
//    {
//      "current": "pnktnjyb996sj4p1523jsklsu324im",
//      "created": "2011-11-07T20:58:34.448Z"
//    }
// Status Codes:
//    201 – Log Created
//    400 – Errors (invalid json, missing or invalid fields, etc)
//    405 - Token 认证失败
func (this *LogController) CreateLog() {

}

// GET /v1/logs/:id
// 根据 ID 查询 Log 日志。
// Example Request:
//    GET /v1/logs/:id HTTP/1.1
//    Host: log.arkors.com
//    X-Arkors-Application-Id: demo
//    X-Arkors-Application-Token: 13b0a8dbddd7c98499a12976ab023ece,1389085779854
//    Accept: application/json
// Example Response:
//    HTTP/1.1 200 OK
//    Vary: Accept
//    Content-Type: application/json
//    {
//      "level": "info",
//      "message": "Create user named genedna.",
//      "root": "pnktnjyb996sj4p156gjtp4im",
//      "up": "pnktnjyb996sj4p156gjtp4im",
//      "current": "pnktnjyb996sj4p156gjtp4im",
//      "created": "2011-11-07T20:58:34.448Z"
//    }
// Status Codes:
//    200 – 返回 Log 日志
//    400 – Errors (invalid json, missing or invalid fields, etc)
//    404 - 没有找到 log 日志
//    405 - Token 认证失败
func (this *LogController) GetLog() {

}
