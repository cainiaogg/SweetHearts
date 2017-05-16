package controllers

import (
	"SweetHearts/models"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpLoadFriendCirclePicture struct {
	beego.Controller
}

func (this *UpLoadFriendCirclePicture) Post() {
	fmt.Println(this.Ctx.Request.MultipartForm)
	Form := this.Ctx.Request.MultipartForm
	fmt.Println(Form.Value)
	username := Form.Value["username"][0]
	id := Form.Value["id"][0]

	f, h, _ := this.GetFile("image")
	pictureName := username + id + h.Filename
	path := filepath.Join(models.FRIEND_CIRCLE_PICTURE, pictureName)
	this.SaveToFile("image", path)

	fmt.Println("upload_friend", id, username, path)
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	FriendCircle := models.FriendCircle{}
	FriendCircle.Id = idInt
	o := orm.NewOrm()
	o.Read(&FriendCircle, "Id")
	FriendCircle.ImagePath += "\t" + models.FRIEND_CIRCLE_IP + pictureName
	o.Update(&FriendCircle)
	f.Close()
	var result []string
	result = append(result, pictureName)
	result = append(result, models.FRIEND_CIRCLE_IP+pictureName)
	this.Data["json"] = result
	this.ServeJSON()
}
