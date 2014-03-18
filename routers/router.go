package routers

import (
	"log/controllers"
	"github.com/astaxie/beego"
)

func init() {
  beego.Router("/", &controllers.MainController{})
}
