package routers

import (
	"SweetHearts/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/user/up_load_headimage", &controllers.UpLoadHeadImage{})
	beego.Router("/user/add_friend", &controllers.AddFriendContoller{})
	beego.Router("/user/get_friend_list", &controllers.GetFriendList{})
	beego.Router("/user/get_add_friend", &controllers.GetAddFriend{})
	beego.Router("/user/friend_yes", &controllers.FriendYes{})
	beego.Router("/pos/update_pos", &controllers.UpdatePosController{})
	beego.Router("/pos/get_pos", &controllers.GetPosController{})
	beego.Router("/news/get_one_note", &controllers.GetOneNote{})
	beego.Router("/circle/get_id", &controllers.GetFriendCircleId{})
	beego.Router("/circle/upload_picture", &controllers.UpLoadFriendCirclePicture{})
	beego.Router("/circle/add_friend_circle", &controllers.AddFriendCirCle{})
	beego.Router("/circle/get_friend_circle", &controllers.GetFriendCircle{})
	beego.Router("/test", &controllers.Test{})
	beego.Router("/chat", &controllers.ChatController{})
	beego.Router("/chat/get_chat_record", &controllers.GetChatRecord{})
}
