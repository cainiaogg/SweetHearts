package controllers

import (
	"SweetHearts/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type GetOneNote struct {
	beego.Controller
}

func GetHash(data string) int {
	ans := 0
	for i := 0; i < len(data); i++ {
		cc, err := strconv.Atoi(fmt.Sprintf("%d", data[i]))
		if err != nil {
			fmt.Println(err)
			continue
		}
		ans += cc - 'a'
	}
	fmt.Println(ans)
	if ans < 0 {
		ans = ans%models.ONE_NOTE_INDEX + models.ONE_NOTE_INDEX
		ans %= models.ONE_NOTE_INDEX
	}
	return ans % models.ONE_NOTE_INDEX
}

func (this *GetOneNote) Get() {
	Username := this.GetString("username")
	user := models.User{}
	user.UserName = Username
	o := orm.NewOrm()
	o.Read(&user, "UserName")
	var hashString string
	if user.Honey > user.UserName {
		hashString = user.Honey + user.UserName
	} else {
		hashString = user.UserName + user.Honey
	}
	fmt.Println(hashString)
	oneNoteSuffix := GetHash(hashString)
	oneNoteSuffix = models.OneNoteIndex - oneNoteSuffix
	fmt.Println(hashString)
	fmt.Println(oneNoteSuffix)

	// f, _ := os.Open("./Content/500")
	// Ans, _ := ioutil.ReadAll(f)
	// this.Ctx.WriteString(string(Ans))
	fmt.Println(models.OneNoteResult[oneNoteSuffix])
	this.TplName = models.OneNoteResult[oneNoteSuffix]
}
