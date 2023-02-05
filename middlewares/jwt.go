package middlewares

import (
	"fmt"
	"go_code/project/byte_douyin/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserClaims struct {
	UserId int64
	jwt.StandardClaims
}

var myKey = []byte("douyin")

// 生成token
func GenerateToken(userLogin models.UserLogin) (string, error) {
	UserClaim := UserClaims{
		UserId:         userLogin.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:		time.Now().Add(7*24*1*time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	//token无效或者token到期
	if err != nil || !claims.Valid {
		return nil, fmt.Errorf("错误的token,err: %v", err)
	}
	return userClaim, nil
}

// 鉴权并得到user_id
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		//用户未登录
		if token == "" {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode:401,
				StatusMsg: "用户未登录",
			})
			c.Abort() //阻止执行
			return
		}
		//验证token
		userClaims, err := AnalyseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode: 401,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		//请求的user_id与token不一致
		userId := c.Query("user_id")
		id, _ := strconv.ParseInt(userId, 10, 64)
		if id != userClaims.UserId {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode: 401,
				StatusMsg:  "请求的user_id与token不一致",
			})
			c.Abort()
			return
		}
		c.Set("user_id", userClaims.UserId)
		c.Next()
	}
}