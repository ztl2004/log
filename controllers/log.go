package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log/models"
)

func init() {
	
	dbHost := beego.AppConfig.String("DBHost")
	dbPort := beego.AppConfig.String("DBPort")
	dbUser := beego.AppConfig.String("DBUser")
	dbPass := beego.AppConfig.String("DBPass")
	dbName := beego.AppConfig.String("DBName")
	
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",dbUser,dbPass,dbHost,dbPort,dbName)

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dbDSN) 

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
	  fmt.Println(err)
	}
}

type LogController struct {
	beego.Controller
}

func (this *LogController) Get() {
	this.Data["Website"] = "test"
	this.Data["Email"] = "test@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *LogController) Post(){
	o := orm.NewOrm()
	o.Using("default")

	log := new(models.Log)
	log.Content = this.GetString("data")

	if log.Content == "" {
		this.Ctx.WriteString("data is empty")
			return
	}

	fmt.Println(o.Insert(log))
	this.TplNames = "index.tpl"

}


