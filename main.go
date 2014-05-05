package main

import (
	"fmt"
	_ "github.com/arkors/log/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("DBUser"),
		beego.AppConfig.String("DBPass"),
		beego.AppConfig.String("DBHost"),
		beego.AppConfig.String("DBPort"),
		beego.AppConfig.String("DBName"))

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbDSN)

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	beego.Run()
}
