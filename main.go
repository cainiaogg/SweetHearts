package main

import (
	"SweetHearts/models"
	_ "SweetHearts/routers"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func addOneNoteIndex() {
	for true {
		time.Sleep(time.Minute * 5)
		models.NowOneNoteTime = time.Now()
		if models.NowOneNoteTime.Sub(models.LastOneNoteTime) > time.Hour*12 {
			models.LastOneNoteTime = models.NowOneNoteTime
			models.OneNoteIndex += 20
			if models.OneNoteIndex > len(models.OneNoteResult) {
				models.OneNoteIndex = len(models.OneNoteResult) - 1
			}
			fmt.Println("OneNoteIndex", models.OneNoteIndex)
		}
	}
}

func init() {
	models.OneNoteIndex = 20
	models.NowOneNoteTime = time.Now()
	models.LastOneNoteTime = time.Now()
	models.RedisClient, _ = cache.NewCache("memory", `{"key":"collectionName","conn":":6379","dbNum":"0"`)
	beego.SetStaticPath("/image", models.IMAGE_PATH)
	beego.SetStaticPath("/static", models.STATIC_PATH)
	beego.SetStaticPath("/chat_front", models.CHAT_STATIC_PATH)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:@/SweetHearts?charset=utf8")
	dirname, _ := ioutil.ReadDir(models.ONE_NOTE_PATH)

	for _, val := range dirname {
		models.OneNoteResult = append(models.OneNoteResult, val.Name())
	}
	go addOneNoteIndex()
}

func main() {
	o := orm.NewOrm()
	o.Using("default")
	orm.RunSyncdb("default", false, true)
	beego.BConfig.Listen.AdminPort = 8088
	beego.Run()
}
