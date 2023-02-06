package middlewares

import (
	"go_code/project/byte_douyin/models"
	"go_code/project/byte_douyin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//JWT中间件,用户权限校验
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
		userClaims, err := utils.AnalyseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.CommonResponse{
				StatusCode: 401,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		//验证请求的user_id和token
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