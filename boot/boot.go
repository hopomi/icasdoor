package boot

import (
	"fmt"
	"icasdoor/models"
	"icasdoor/services/passwd"
	"icasdoor/services/rsa"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
)

func genDefaultRsa() {
	var err error
	pri, pub := rsa.GenRsa()
	err = os.MkdirAll("files/rsa", 0777)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("files/rsa/rsa", pri, 0777)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("files/rsa/rsa.pub", pub, 0777)
	if err != nil {
		panic(err)
	}
}

func OrmSyncRegister() {
	orm.RegisterDataBase("default", "postgres", config.DefaultString("pgsqlconn.default.link", ""))
	models.OrmRegister()
	orm.RunSyncdb("default", false, true)
}

func CreatDefaultSuperUser() {
	o := orm.NewOrm()
	passwd, err := passwd.GenPasswordMD5("icasdoor")
	user := models.Users{Username: "icasdoor", Password: passwd, IsSupper: true, IsStuff: true}
	if err != nil {
		panic(err)
	}
	created, id, err := o.ReadOrCreate(&user, "Username")
	if err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	}
	profile := models.Profile{Users: &user}
	if created, id, err := o.ReadOrCreate(&profile, "users_id"); err == nil {
		if created {
			fmt.Println("New Insert an object. Id:", id)
		} else {
			fmt.Println("Get an object. Id:", id)
		}
	}

}

func Boot() {
	// genDefaultRsa()
	OrmSyncRegister()
	CreatDefaultSuperUser()
}
