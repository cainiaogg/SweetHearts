package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
)

type GetAddFriend struct {
	beego.Controller
}

func (this *GetAddFriend) Get() {
	Username := this.GetString("username")
	var result []string
	result = models.RedisClient.Get(Username).([]string)
	this.Data["json"] = result
	fmt.Println(result)
	this.ServeJSON()
}
