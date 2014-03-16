package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/log?charset=utf8")
}

type LogController struct {
	beego.Controller
}

func (this *LogController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *LogController) Post(){
	o := orm.NewOrm()
	o.Using("default")

	log := new(models.Log)
	log.Content = this.GetString("log")

	fmt.Println(o.Insert(log))

}


