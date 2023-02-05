package handlers

import (
	"go_code/project/byte_douyin/models"
	"go_code/project/byte_douyin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserLoginResponse struct {
	models.CommonResponse
	*services.LoginResponse
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	LoginResponse, err := services.PostUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, UserLoginResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 500,
				StatusMsg: err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		LoginResponse: LoginResponse,
	})
	
}