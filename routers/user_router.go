package routers

import (
	"go_code/project/byte_douyin/handlers"
	"go_code/project/byte_douyin/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouters(g *gin.RouterGroup) {
	r := g.Group("/user")
	{	
		r.GET("/", middlewares.JWTMiddleware(), handlers.UserInfoHandler)
		r.POST("/register/", middlewares.PwdHashMiddleWare(), handlers.UserRegisterHandler)
		r.POST("/login/", handlers.UserLoginHandler)
	}
}