package controllers

import (
	"SweetHearts/models"
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

type UpLoadHeadImage struct {
	beego.Controller
}

func (this *UpLoadHeadImage) Post() {
	// f_out, err := os.Create("test.png")
	// if err != nil {
	// fmt.Println(err)
	// }
	// size, err := f_out.Write(this.Ctx.Input.RequestBody)
	// if err != nil {
	// fmt.Println(err)
	// fmt.Println(size)
	// }
	f, h, _ := this.GetFile("image") //获取上传的文件
	// username := this.GetString("username")
	path := filepath.Join(models.UPLOAD_PATH, h.Filename)
	// this.Ctx.WriteString(path)
	f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err := os.Mkdir(models.UPLOAD_PATH, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	this.SaveToFile("image", path)
	fmt.Println(h.Filename)
	this.Ctx.WriteString(h.Filename)
}
