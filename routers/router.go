package routers

import (
	"github.com/arkors/log/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Create Log
	beego.Router("/v1/logs", &controllers.LogController{}, "post:CreateLog")
	//Get Log
	beego.Router("/v1/logs/:id", &controllers.LogController{}, "get:GetLog")
	//Search Log
	beego.Router("/v1/search", &controllers.SearchController{}, "post:SearchLog")
	//Default
	beego.Router("/", &controllers.MainController{})
}
