package main

import (
	_ "log/routers"
	"log/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/log",&controllers.LogController{})
	beego.Run()
}

