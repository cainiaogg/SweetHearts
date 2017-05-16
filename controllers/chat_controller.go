package controllers

import (
	"SweetHearts/models"
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gorilla/websocket"
)

type ChatController struct {
	beego.Controller
}

func (this *ChatController) Get() {
	UserName := this.GetString("username")
	FriendName := this.GetString("friendname")
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	var RoomName string
	if UserName > FriendName {
		RoomName = UserName + FriendName
	} else {
		RoomName = FriendName + UserName
	}
	defer func() {
		fmt.Println("connect end..")
		models.ChatRoomMap[RoomName].Leave(UserName)
		if len(models.ChatRoomMap[RoomName].ChatUserMap) == 0 {
			delete(models.ChatRoomMap, RoomName)
		}
		this.Ctx.WriteString("connect end")
	}()
	if _, ok := models.ChatRoomMap[RoomName]; ok != true {
		models.ChatRoomMap[RoomName] = models.NewChatRoom()
	}
	models.ChatRoomMap[RoomName].Join(UserName, ws)
	for {
		_, p, err := ws.ReadMessage()
		if len(string(p)) == 0 {
			continue
		}
		o := orm.NewOrm()
		chatRecord := models.ChatRecord{}
		chatRecord.From = UserName
		chatRecord.To = FriendName
		chatRecord.Content = string(p)
		chatRecord.TimeStamp = time.Now().UnixNano() / 1000000
		fmt.Println("insert chatRecord")
		o.Insert(&chatRecord)
		models.ChatRoomMap[RoomName].SendMessage(UserName, string(p), time.Now().UnixNano()/1000000)
		fmt.Println(string(p))
		if err != nil {
			fmt.Println("err " + err.Error())
			this.Ctx.WriteString("JIJI")
			return
		}
	}
}
