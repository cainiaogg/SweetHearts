package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	simplejson "github.com/bitly/go-simplejson"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	var logInformation map[string]interface{}
	logInformation = make(map[string]interface{})
	fmt.Println(string(this.Ctx.Input.RequestBody))
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	logInformation, err = js.Map()
	if err != nil {
		beego.Informational(err.Error())
	}

	username := logInformation["username"]
	password := logInformation["password"]
	o := orm.NewOrm()
	user := models.User{UserName: username.(string), PassWord: password.(string)}
	err = o.Read(&user, "UserName", "PassWord")

	errorInformation := new(models.ErrorInformation)
	if err == orm.ErrNoRows {
		errorInformation.ErrorNumber = 0
		errorInformation.ErrorString = "用户名密码错误"
	} else {
		errorInformation.ErrorNumber = 1
		errorInformation.ErrorString = user.Honey
	}
	this.Data["json"] = &errorInformation
	this.ServeJSON()
	// this.Ctx.WriteString("sb")
}
