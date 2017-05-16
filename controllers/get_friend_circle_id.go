package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetFriendCircleId struct {
	beego.Controller
}

func (this *GetFriendCircleId) Get() {
	friendCircle := models.FriendCircle{}
	o := orm.NewOrm()
	o.Insert(&friendCircle)
	this.Ctx.WriteString(fmt.Sprintf("%d", friendCircle.Id))
}
