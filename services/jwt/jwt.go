package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"icasdoor/models"
	"math/rand"
	"os"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/golang-jwt/jwt/v5"
)

var pubKey, _ = os.ReadFile("files/rsa/rsa.pub")
var priKey, _ = os.ReadFile("files/rsa/rsa")

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}
func parsePriKeyBytes(buf []byte) (*rsa.PrivateKey, error) {
	p := &pem.Block{}
	p, buf = pem.Decode(buf)
	if p == nil {
		return nil, errors.New("parse key error")
	}
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}

func parsePubKeyBytes(pub_key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pub_key)
	if block == nil {
		return nil, errors.New("block nil")
	}
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func ValidJwt(token_string string) (*models.CommonClaims, error) {
	token, err := jwt.ParseWithClaims(token_string, &models.CommonClaims{}, func(token *jwt.Token) (interface{}, error) {
		pub, err := parsePubKeyBytes(pubKey)
		if err != nil {
			logs.Error(err)
			panic(err)
		}
		return pub, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("claim invalid")
	}

	claims, ok := token.Claims.(*models.CommonClaims)
	if !ok {
		return nil, errors.New("invalid claim type")
	}

	return claims, nil
}

func GenJwt() (string, error) {
	claim := models.CommonClaims{
		UserID:     000001,
		Username:   "Tom",
		GrantScope: "read_user_info",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Auth_Server",                                      // 签发者
			Subject:   "Tom",                                              // 签发对象
			Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},         //签发受众
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 10)), //过期时间
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),    //最早使用时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     //签发时间
			ID:        randStr(10),                                        // jwt ID, 类似于盐值
		},
	}
	rsa_pri_key, err := parsePriKeyBytes(priKey)
	if err != nil {
		logs.Error(err)
		panic(err)
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claim).SignedString(rsa_pri_key)
	return token, err
}
