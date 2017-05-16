package controllers

import (
	"SweetHearts/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetChatRecord struct {
	beego.Controller
}

func (this *GetChatRecord) Get() {
	UserName := this.GetString("username")
	FriendName := this.GetString("friendname")
	TimeStamp, err := this.GetInt64("timestamp")
	if err != nil {
		fmt.Println("GetChatRecord time err" + err.Error())
	}
	var chatRecordList []models.ChatRecord
	sqlString := fmt.Sprintf("SELECT chat_record.* from chat_record where chat_record.time_stamp < %v and ((chat_record.from=\"%s\" and chat_record.to=\"%s\") or (chat_record.from=\"%s\" and chat_record.to=\"%s\")) order by chat_record.time_stamp DESC limit 10", TimeStamp, UserName, FriendName, FriendName, UserName)
	o := orm.NewOrm()
	o.Raw(sqlString).QueryRows(&chatRecordList)
	fmt.Println("get chat record sql: " + sqlString)
	var ans []interface{}
	for _, val := range chatRecordList {
		var ans1 []interface{}
		if val.From == UserName {
			ans1 = append(ans1, 1)
			ans1 = append(ans1, val)
			ans = append(ans, ans1)
			continue
		}
		ans1 = append(ans1, 2)
		ans1 = append(ans1, val)
		ans = append(ans, ans1)
	}
	this.Data["json"] = ans
	this.ServeJSON()
}
