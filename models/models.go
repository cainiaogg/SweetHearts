package models

import (
	"time"

	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
)

type ErrorInformation struct {
	ErrorNumber int64
	ErrorString interface{}
}

type ChatRecord struct {
	Id        int
	From      string
	To        string
	Content   string
	TimeStamp int64
}

type FriendCircle struct {
	Id        int
	UserName  string
	HeadImage string
	Content   string
	TimeStamp int64
	ImagePath string
	IsHoney   string
}

type User struct {
	Id       int
	UserName string `orm:"unique"`
	PassWord string
	Profile  *Profile `orm:"rel(one)"`
	Honey    string
}

type Friend struct {
	Id         int
	FriendName string
	UserName   string
}

type Profile struct {
	Id        int
	RealName  string
	Sex       string
	User      *User `orm:"reverse(one)"`
	HeadImage string
}

func NewProfile() *Profile {
	var result = new(Profile)
	result.RealName = ""
	result.Sex = ""
	return result
}

var OneNoteResult []string
var OneNoteIndex int
var LastOneNoteTime time.Time
var NowOneNoteTime time.Time
var ChatRoomMap map[string]*ChatRoom

var RedisClient cache.Cache

const REDIS_ADDFRIEND_TIME_OUT time.Duration = time.Second * 1000000000

const GET_FRIEND_CIRCLE_PAGE = 10

const ORIGIN_HEAD_IMAGE string = UPLOAD_PATH + "/origin.jpg"

const FRIEND_CIRCLE_IP string = "http://118.89.200.173/image/friend_circle_picture/"

const GET_HEAD_IMAGE string = "http://118.89.200.173/image/head_image/"

const PROJECT_PATH string = "/usr/local/Cellar/go/1.6.2/src/SweetHearts"

const UPLOAD_PATH string = PROJECT_PATH + "/image/head_image"

const STATIC_PATH string = PROJECT_PATH + "/static"

const FRIEND_CIRCLE_PICTURE string = PROJECT_PATH + "/image/friend_circle_picture"

const IMAGE_PATH string = PROJECT_PATH + "/image"

const ONE_NOTE_PATH string = PROJECT_PATH + "/views"

const ONE_NOTE_INDEX int = 20

func init() {
	ChatRoomMap = make(map[string]*ChatRoom)
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Profile))
	orm.RegisterModel(new(Friend))
	orm.RegisterModel(new(FriendCircle))
	orm.RegisterModel(new(ChatRecord))
}
