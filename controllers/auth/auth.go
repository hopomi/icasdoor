package auth

import (
	"context"
	"icasdoor/api/payload"
	"icasdoor/boot"
	"icasdoor/services/auth"
	"icasdoor/services/jwt"

	beego "github.com/beego/beego/v2/server/web"
)

// token 处理相关。
type AuthController struct {
	beego.Controller
}

func (o *AuthController) tokenInBlackList(token string) bool {
	return boot.DefaultRedisClient.SIsMember(context.Background(), "icasdoor-jwt/blacklist", token).Val()
}

// @Title GenJwt
// @Description 登录、验证
// @Param	body		body 	payload.GenJwtReq	true	"The object content"
// @Success 200 {object} payload.GenJwtResp
// @router /jwt/gen [post]
func (o *AuthController) GenJwt() {
	var err error
	var payload payload.GenJwtReq
	o.BindJSON(&payload)
	authService := new(auth.AuthService)
	valid, err := authService.ValidUsernameAndPasswd(payload.Username, payload.Password)
	if !valid || err != nil {
		auth.UnAuthentication403(&o.Controller, "username or password invalid")
		return
	}
	user, err := authService.GetUser()
	if err != nil {
		auth.UnAuthentication403(&o.Controller, "username or password invalid")
		o.ServeJSON()
		return
	}
	s, err := jwt.GenJwt(int(user.Id), user.Username)
	if err != nil {
		panic(err)
	}
	o.Data["json"] = s
	o.ServeJSON()
}

// @Title ValidJwt
// @Description 验证、解码
// @Param	body		body 	payload.ValidJwtReq	true	"The object content"
// @Success 200 {object} payload.ValidJwtReq
// @router /jwt/valid [post]
func (o *AuthController) ValidJwt() {
	var err error
	var req payload.ValidJwtReq
	err = o.BindJSON(&req)
	if err != nil {
		panic(err)
	}
	if o.tokenInBlackList(req.Token) {
		auth.UnAuthentication403(&o.Controller, "invalid, token in black list")
		return
	}
	o.Data["json"], err = jwt.ValidJwt(req.Token)
	if err != nil {
		panic(err)
	}
	o.ServeJSON()
}

// @Title RefreshJwt
// @Description 刷新
// @Param	body		body 	payload.ValidJwtReq	true	"The object content"
// @Success 200 {object} payload.ValidJwtReq
// @router /jwt/refresh [post]
func (o *AuthController) RefreshJwt() {
	var err error
	var req payload.ValidJwtReq
	err = o.BindJSON(&req)
	if err != nil {
		panic(err)
	}
	claims, err := jwt.ValidJwt(req.Token)
	if err != nil {
		panic(err)
	}
	if o.tokenInBlackList(req.Token) {
		auth.UnAuthentication403(&o.Controller, "invalid, token in black list")
		return
	}
	s, err := jwt.GenJwt(claims.UserID, claims.Username)
	if err != nil {
		panic(err)
	}
	c := boot.DefaultRedisClient.SAdd(context.Background(), "icasdoor-jwt/blacklist", req.Token).Val()
	if c == 0 {
		panic("")
	}
	o.Data["json"] = s
	o.ServeJSON()
}
