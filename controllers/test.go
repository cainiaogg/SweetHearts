package controllers

import (
	"SweetHearts/models"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Test struct {
	beego.Controller
}

func (this *Test) Post() {
	fmt.Println(this.Ctx.Request.Body)
	fmt.Println(this.Ctx.Request.MultipartForm)
	f, h, _ := this.GetFile("image")
	username := "jiba"
	id := "1"
	picture_name := username + id + h.Filename
	path := filepath.Join(models.FRIEND_CIRCLE_PICTURE, picture_name)
	this.SaveToFile("image", path)

	fmt.Println("upload_friend", id, username, path)
	id_int, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	FriendCircle := models.FriendCircle{}
	FriendCircle.Id = id_int
	o := orm.NewOrm()
	o.Read(&FriendCircle, "Id")
	FriendCircle.ImagePath += "\t" + models.FRIEND_CIRCLE_IP + picture_name
	o.Update(&FriendCircle)
	f.Close()
	this.Ctx.WriteString(picture_name)
}
