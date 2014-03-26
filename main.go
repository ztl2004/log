package main

import (
	"fmt"
	_ "log/routers"
	"log/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
		)

func init() {

	dbHost := beego.AppConfig.String("DBHost")
	dbPort := beego.AppConfig.String("DBPort")
	dbUser := beego.AppConfig.String("DBUser")
	dbPass :=beego.AppConfig.String("DBPass")
	dbName :=beego.AppConfig.String("DBName")

	dbDSN:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",dbUser,dbPass,dbHost,dbPort,dbName)

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbDSN)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Router("/log",&controllers.LogController{})
		beego.Run()
}

