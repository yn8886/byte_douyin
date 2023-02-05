package middlewares

import (
	"errors"
	"go_code/project/byte_douyin/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PwdHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func PwdHashMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Query("password")
		//检验密码是否规范
		if regexp.MustCompile(`^.{6,32}$`).FindString(password) == "" {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode: 500,
				StatusMsg: errors.New("密码长度超出限制").Error(),
			})
			c.Abort()
			return
		}
		//密码加密
		hash, err := PwdHash(password)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode: 500,
				StatusMsg: err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("hash_password", hash)
		c.Next()
	}
}

