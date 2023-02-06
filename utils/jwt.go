package utils

import (
	"fmt"
	"go_code/project/byte_douyin/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId int64
	jwt.StandardClaims
}

var myKey = []byte("douyin")

//生成token
func GenerateToken(userLogin models.UserLogin) (string, error) {
	UserClaim := UserClaims{
		UserId:         userLogin.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:		time.Now().Add(7*24*time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	//token无效或者token到期
	if err != nil || !claims.Valid {
		return nil, fmt.Errorf("token错误: %v", err)
	}
	return userClaim, nil
}
