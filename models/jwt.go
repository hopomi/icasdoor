package models

import "github.com/golang-jwt/jwt/v5"

type CommonClaims struct {
	UserID     int    `json:"user_id"`
	Username   string `json:"username"`
	GrantScope string `json:"grant_scope"`
	jwt.RegisteredClaims
}
