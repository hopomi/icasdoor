package auth

import (
	"context"
	"icasdoor/api/payload"
	"icasdoor/boot"
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
// @Description redis get
// @Param	body		body 	payload.GenJwtReq	true	"The object content"
// @Success 200 {object} payload.GenJwtResp
// @router /jwt/gen [post]
func (o *AuthController) GenJwt() {
	var err error
	s, err := jwt.GenJwt()
	if err != nil {
		panic(err)
	}
	o.Data["json"] = s
	o.ServeJSON()
}

// @Title ValidJwt
// @Description redis get
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
		o.Ctx.Output.SetStatus(403)
		o.Data["json"] = map[string]interface{}{"message": "invalid, token in black list"}
		o.ServeJSON()
		return
	}
	o.Data["json"], err = jwt.ValidJwt(req.Token)
	if err != nil {
		panic(err)
	}
	o.ServeJSON()
}

// @Title RefreshJwt
// @Description redis get
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
	_, err = jwt.ValidJwt(req.Token)
	if err != nil {
		panic(err)
	}
	if o.tokenInBlackList(req.Token) {
		o.Ctx.Output.SetStatus(403)
		o.Data["json"] = map[string]interface{}{"message": "invalid, token in black list"}
		o.ServeJSON()
		return
	}
	s, err := jwt.GenJwt()
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
