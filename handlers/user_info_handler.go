package handlers

import (
	"go_code/project/byte_douyin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	models.CommonResponse
	User *models.User `json:"user"`
}

func UserInfoHandler(c *gin.Context) {
	val, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, UserResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 500,
				StatusMsg: "解析user_id出错",
			},
		})
	}
	id := val.(int64)
	user := models.User{}
	err := models.QueryUserInfoByUserId(id, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 500,
				StatusMsg: err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		User: &user,
	})
}

