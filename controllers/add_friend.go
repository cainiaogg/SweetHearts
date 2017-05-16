package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	simplejson "github.com/bitly/go-simplejson"
)

type AddFriendContoller struct {
	beego.Controller
}

func (this *AddFriendContoller) Post() {
	var form map[string]interface{}
	form = make(map[string]interface{})
	errorInformation := new(models.ErrorInformation)
	errorInformation.ErrorNumber = 1
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	form, err = js.Map()
	o := orm.NewOrm()
	user := models.User{}
	user.UserName = form["to"].(string)
	err = o.Read(&user, "UserName")

	if err != nil {
		errorInformation.ErrorNumber = 0
		beego.Error(err.Error())
	}

	from := form["from"]
	to := form["to"]
	isHoney := form["isHoney"]
	var result []string
	result_pre := models.RedisClient.Get(to.(string))
	switch v := result_pre.(type) {
	case []string:
		result = v
	}
	fmt.Println(models.REDIS_ADDFRIEND_TIME_OUT)
	if result == nil {
		tmp := []string{}
		tmp = append(tmp, from.(string)+isHoney.(string))
		models.RedisClient.Put(to.(string), tmp, models.REDIS_ADDFRIEND_TIME_OUT)
	} else {
		result = append(result, from.(string)+isHoney.(string))
		models.RedisClient.Put(to.(string), result, models.REDIS_ADDFRIEND_TIME_OUT)
	}
	fmt.Println(models.RedisClient.Get(to.(string)))
	this.Data["json"] = &errorInformation
	this.ServeJSON()
}
