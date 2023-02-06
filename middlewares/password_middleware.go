package middlewares

import (
	"errors"
	"go_code/project/byte_douyin/models"
	"go_code/project/byte_douyin/utils"
	"net/http"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func PwdHashMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.Query("password")
		//验证密码规范性
		if utf8.RuneCountInString(password) > 32 {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode: 500,
				StatusMsg: errors.New("密码长度超出限制").Error(),
			})
			c.Abort()
			return
		}
		//密码加密
		hash, err := utils.PwdHash(password)
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

