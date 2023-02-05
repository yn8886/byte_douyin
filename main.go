package main

import (
	"go_code/project/byte_douyin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.BaseRouter(r)
	r.Run(":8080")
}