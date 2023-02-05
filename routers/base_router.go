package routers

import "github.com/gin-gonic/gin"

func BaseRouter(g *gin.Engine) {
	br := g.Group("/douyin")
	UserRouters(br)
}