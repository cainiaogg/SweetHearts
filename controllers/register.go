package controllers

import (
	"SweetHearts/models"
	"io"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	simplejson "github.com/bitly/go-simplejson"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Post() {
	var registerInformation map[string]interface{}
	registerInformation = make(map[string]interface{})
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	registerInformation, err = js.Map()
	if err != nil {
		beego.Informational(err.Error())
	}
	username := registerInformation["username"]
	password := registerInformation["password"]
	user := new(models.User)
	user.UserName = username.(string)
	user.PassWord = password.(string)
	profile := new(models.Profile)
	user.Profile = profile
	o := orm.NewOrm()
	errorInformation := new(models.ErrorInformation)
	errorInformation.ErrorNumber, errorInformation.ErrorString = o.Insert(user)
	if errorInformation.ErrorString != nil {
		this.Data["json"] = &errorInformation
		this.ServeJSON()
		return
	}
	errorInformation.ErrorNumber, errorInformation.ErrorString = o.Insert(profile)
	user.Profile = profile
	o.Update(user)
	this.Data["json"] = &errorInformation
	distName := models.UPLOAD_PATH + "/" + username.(string) + ".jpg"
	distFile, _ := os.Create(distName)
	srcFile, _ := os.Open(models.ORIGIN_HEAD_IMAGE)
	io.Copy(distFile, srcFile)
	defer distFile.Close()
	defer srcFile.Close()
	this.ServeJSON()
}
