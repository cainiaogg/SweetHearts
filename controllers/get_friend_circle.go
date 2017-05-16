package controllers

import (
	"SweetHearts/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetFriendCircle struct {
	beego.Controller
}

func (this *GetFriendCircle) Get() {
	username := this.GetString("username")
	page, _ := this.GetInt("page")
	intent := models.GET_FRIEND_CIRCLE_PAGE
	userTmp := models.User{}
	userTmp.UserName = username

	var friend []models.FriendCircle
	sqlString := fmt.Sprintf("SELECT distinct friend_circle.* FROM friend, friend_circle where (friend.user_name =\"%s\" and SweetHearts.friend.friend_name=SweetHearts.friend_circle.user_name) or (friend_circle.user_name=\"%s\") order by friend_circle.time_stamp DESC", username, username)
	o := orm.NewOrm()
	o.Read(&userTmp, "UserName")
	honeyName := userTmp.Honey

	var result []models.FriendCircle
	num64, _ := o.Raw(sqlString).QueryRows(&friend)
	fmt.Printf("friend: \n%v\n", friend)
	begin := (page - 1) * intent
	end := page*intent - 1
	num := int(num64)
	if begin >= num {
		var cc []string
		this.Data["json"] = cc
		this.ServeJSON()
		return
	}
	if begin >= num {
		begin = num - 1
	}
	if end >= num {
		end = num - 1
	}
	fmt.Println(begin, end)
	for index, val := range friend {
		if index < begin || index > end {
			continue
		}
		if val.IsHoney == "0" {
			result = append(result, val)
			continue
		}
		if val.UserName == username {
			result = append(result, val)
			continue
		}
		if honeyName == "" {
			if val.UserName == username {
				result = append(result, val)
			}
			continue
		}
		if honeyName != val.UserName {
			continue
		}
		result = append(result, val)
	}
	var resultList []interface{}
	for _, val := range result {
		var resultMap map[string]interface{}
		resultMap = make(map[string]interface{})
		resultMap["Id"] = val.Id
		resultMap["UserName"] = val.UserName
		resultMap["HeadImage"] = models.GET_HEAD_IMAGE + val.UserName + ".jpg"
		resultMap["Content"] = val.Content
		resultMap["TimeStamp"] = val.TimeStamp
		if val.ImagePath != "" {
			resultMap["ImagePath"] = strings.Split(val.ImagePath, "\t")[1:]
		} else {
			resultMap["ImagePath"] = val.ImagePath
		}
		resultMap["IsHoney"] = val.IsHoney
		resultList = append(resultList, resultMap)
	}
	// this.Data["json"] = resultList
	// this.Data["json"] = result
	this.Data["json"] = resultList
	fmt.Println(resultList)
	this.ServeJSON()
}
