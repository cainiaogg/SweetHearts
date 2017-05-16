package controllers

import (
	"SweetHearts/models"

	"github.com/astaxie/beego"
)

type GetPosController struct {
	beego.Controller
}

func (this *GetPosController) Get() {
	Username := this.GetString("username")
	var result []string
	result = models.RedisClient.Get(Username + "pos").([]string)
	this.Data["json"] = result
	this.ServeJSON()
}
