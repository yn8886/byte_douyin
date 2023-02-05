package handlers

import (
	"go_code/project/byte_douyin/models"
	"go_code/project/byte_douyin/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserRegisterResponse struct {
	models.CommonResponse
	*services.RegisterResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	val, _ := c.Get("hash_password")
	password := val.(string)
	registerResponse, err := services.PostUserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserRegisterResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 500,
				StatusMsg: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, UserRegisterResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		RegisterResponse: registerResponse,
	})
}