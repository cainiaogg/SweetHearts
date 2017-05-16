package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
)

type UpdatePosController struct {
	beego.Controller
}

func (this *UpdatePosController) Get() {
	Username := this.GetString("username")
	Lat := this.GetString("lat")
	Long := this.GetString("long")
	TimeStamp := this.GetString("timestamp")
	var result []string
	result = append(result, Username)
	result = append(result, Lat)
	result = append(result, Long)
	result = append(result, TimeStamp)
	models.RedisClient.Put(Username+"pos", result, models.REDIS_ADDFRIEND_TIME_OUT)
	fmt.Println(models.RedisClient.Get(Username + "pos"))
	this.Ctx.WriteString("1")
}
