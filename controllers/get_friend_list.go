package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetFriendList struct {
	beego.Controller
}

func (this *GetFriendList) Get() {
	var friend []*models.Friend
	Username := this.GetString("username")

	orm.NewOrm().QueryTable("Friend").Filter("UserName", Username).RelatedSel().All(&friend)
	fmt.Println(*friend[0])
	friendList := []string{}
	for _, v := range friend {
		friendList = append(friendList, v.FriendName)
	}
	this.Data["json"] = friendList
	this.ServeJSON()
}
