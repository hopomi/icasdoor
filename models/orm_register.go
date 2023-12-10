package models

import "github.com/beego/beego/v2/client/orm"

func OrmRegister() {
	orm.RegisterModelWithPrefix("user_", new(Profile), new(Users))
}
