package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"SweetHearts/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	simplejson "github.com/bitly/go-simplejson"
)

type AddFriendCirCle struct {
	beego.Controller
}

func (this *AddFriendCirCle) Post() {
	var registerInformation = make(map[string]interface{})
	js, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	// fmt.Println("circle", this.Ctx.Input.RequestBody)
	registerInformation, err = js.Map()
	fmt.Printf("circle\n%+v\n ", registerInformation)
	if err != nil {
		beego.Informational(err.Error())
	}
	username := registerInformation["username"].(string)
	content := registerInformation["content"].(string)
	id := registerInformation["id"].(string)
	isHoney_pre := registerInformation["ishoney"].(json.Number)
	isHoney := string(isHoney_pre)
	if isHoney == "0" {
		isHoney = "1"
	} else {
		isHoney = "0"
	}
	id_int, err := strconv.Atoi(id)
	currentTime := time.Now().Unix()
	friend := models.FriendCircle{}
	friend.Id = id_int
	o := orm.NewOrm()
	o.Read(&friend, "Id")
	friend.TimeStamp = currentTime
	friend.Content = content
	friend.UserName = username
	friend.IsHoney = isHoney
	fmt.Printf("%+v\n", friend)
	o.Update(&friend)
	this.Ctx.WriteString("1")
}
