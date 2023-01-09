package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type HmacUser struct {
	Id       string `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
}

type MyClaims struct {
	Id       string `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtKey = []byte("abc")

func GenerateToken(u HmacUser) (string, error) {
	// 定义过期时间,7天后过期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &MyClaims{
		Id:       u.Id,
		Username: u.Username,
		Uuid:     u.Uuid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发布时间
			Subject:   "token",               // 主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
