package controllers

import (
	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

// POST /v1/search
// 根据 ID 查询上下游所有的信息
// Example Request:
//    POST /v1/search HTTP/1.1
//    Host: log.arkors.com
//    X-Arkors-Application-Id: demo
//    X-Arkors-Application-Token: 13b0a8dbddd7c98499a12976ab023ece,1389085779854
//    Accept: application/json
//    {
//      "query": "pnktnjyb996sj4p156gjtp4im",
//    }
// Example Response:
//    HTTP/1.1 201 OK
//    Vary: Accept
//    Location: https://auth.arkors.com/v1/users/132
//    Content-Type: application/json
//    {
//      "query":"pnktnjyb996sj4p156gjtp4im",
//      "results": 12,
//      "results" : [
//          {"level": "info", "message": "Create user named genedna.", "root": "pnktnjyb996sj4p156gjtp4im", "up": "pnktnjyb996sj4p156gjtp4im", "current": "pnktnjyb996sj4p156gjtp4im", "created": "2011-11-07T20:58:34.448Z"}
//      ]
//    }
// Status Codes:
//    200 – 返回日志
//    400 – Errors (invalid json, missing or invalid fields, etc)
//    405 - Token 认证失败
func (this *SearchController) SearchLog() {

}
