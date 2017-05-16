package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type FriendYes struct {
	beego.Controller
}

func (this *FriendYes) Get() {
	Username := this.GetString("username")
	FriendName := this.GetString("friendname")
	IsHoney := this.GetString("isHoney")
	Agree := this.GetString("agree")
	Add := FriendName + IsHoney
	var result []string
	var result_new []string
	err := models.RedisClient.Get(Username)
	if err == nil {
		this.Ctx.WriteString("0")
		return
	}
	result = err.([]string)
	for _, val := range result {
		fmt.Println(val, Add, val == Add)
		if val == Add {
			continue
		}
		result_new = append(result_new, val)
	}
	fmt.Println(result)
	fmt.Println(result_new)
	models.RedisClient.Put(Username, result_new, models.REDIS_ADDFRIEND_TIME_OUT)
	if Agree == "0" {
		this.Ctx.WriteString("1")
		return
	}
	friend := models.Friend{}
	friend.FriendName = FriendName
	friend.UserName = Username
	friend1 := models.Friend{}
	friend1.FriendName = Username
	friend1.UserName = FriendName
	o := orm.NewOrm()
	o.Insert(&friend)
	o.Insert(&friend1)
	if IsHoney == "0" {
		this.Ctx.WriteString("1")
		return
	}
	user := models.User{}
	user.UserName = Username
	o.Read(&user, "UserName")
	user.Honey = FriendName
	o.Update(&user)
	user1 := models.User{}
	user1.UserName = FriendName
	o.Read(&user1, "UserName")
	user1.Honey = Username
	o.Update(&user1)
	this.Ctx.WriteString("1")
}
