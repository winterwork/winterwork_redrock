package main

import (
	"douban/controller"
	"douban/middleware"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	tool.OpenDb()

	r := gin.Default()

	r.Use(middleware.Cors())

	registerRouter(r)

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}

//注册路由
func registerRouter(r *gin.Engine) {
	controller.UserRouter(r)
	controller.AdminRouter(r)
	controller.MessageRouter(r)
	controller.MovieRouter(r)
	controller.MemberRouter(r)
}
