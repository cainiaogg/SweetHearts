package main

import (
	"SweetHearts/models"
	_ "SweetHearts/routers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func S2B(s *string) []byte {
	return *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(s))))
}

func addOneNoteIndex() {
	for true {
		time.Sleep(time.Minute * 5)
		models.NowOneNoteTime = time.Now()
		if models.NowOneNoteTime.Sub(models.LastOneNoteTime) > time.Hour*12 {
			models.LastOneNoteTime = models.NowOneNoteTime
			models.OneNoteIndex += 20
			err := models.RedisClient.Put("LastOneNoteTime", models.LastOneNoteTime, models.REDIS_ADDFRIEND_TIME_OUT)
			err = models.RedisClient.Put("OneNoteIndex", models.OneNoteIndex, models.REDIS_ADDFRIEND_TIME_OUT)
			if err != nil {
				fmt.Println("error update OneNoteIndex err:", err)
			}
			if models.OneNoteIndex > len(models.OneNoteResult) {
				models.OneNoteIndex = len(models.OneNoteResult) - 1
			}
			fmt.Println("OneNoteIndex", models.OneNoteIndex)
		}
	}
}

func init() {
	models.M_REDIS = new(sync.Mutex)
	models.RedisClient = models.NewMyRedisClient()
	models.RedisClient = models.NewMyRedisClient()
	if ans_one_note_index := models.RedisClient.Get("OneNoteIndex"); ans_one_note_index != nil {
		models.OneNoteIndex, _ = strconv.Atoi(ans_one_note_index.(string))
	} else {
		models.OneNoteIndex = 20
	}
	models.NowOneNoteTime = time.Now()
	if ans_last_one_time := models.RedisClient.Get("LastOneNoteTime"); ans_last_one_time != nil {
		string_ans_last_one_time := ans_last_one_time.(string)
		byte_time := S2B(&string_ans_last_one_time)
		json.Unmarshal(byte_time, &models.LastOneNoteTime)
	} else {
		models.LastOneNoteTime = time.Now()
	}
	// models.RedisClient, _ = cache.NewCache("memory", `{"key":"collectionName","conn":":6379","dbNum":"0"`)
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
