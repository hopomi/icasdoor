package auth

import (
	"icasdoor/models"
	"icasdoor/services/passwd"

	"github.com/beego/beego/v2/client/orm"

	beego "github.com/beego/beego/v2/server/web"
)

type AuthService struct {
	user models.Users
}

func (s *AuthService) ValidUsernameAndPasswd(username, inPasswd string) (bool, error) {
	inPasswd, err := passwd.GenPasswordMD5(inPasswd)
	if err != nil {
		return false, err
	}
	var user = models.Users{Username: username, Password: inPasswd, Activate: true}
	o := orm.NewOrm()
	err = o.Read(&user, "username", "password", "activate")
	if err != nil {
		if err == orm.ErrNoRows {
			return false, err
		} else if err == orm.ErrMissPK {
			return false, err
		}
		return false, err
	}
	s.user = user
	return true, nil
}

func (s *AuthService) GetUser() (models.Users, error) {
	return s.user, nil
}

func UnAuthentication403(o *beego.Controller, messages string) {
	o.Ctx.Output.SetStatus(403)
	if len(messages) <= 0 {
		messages = "UnAuthentication"
	}
	o.Data["json"] = map[string]interface{}{"message": messages}
	o.ServeJSON()
}
