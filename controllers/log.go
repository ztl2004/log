package controllers

import (
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
  "log/models"
)

type LogController struct {
  beego.Controller
}

func (this *LogController) Get() {
  this.Data["Website"] = "test"
  this.Data["Email"] = "test@gmail.com"
  this.TplNames = "index.tpl"
}

func (this *LogController) Post() {
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
