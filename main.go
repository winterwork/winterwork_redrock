package main

import (
	"douban/controller"
	"douban/middleware"
	"douban/tool"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	tool.OpenDb()

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	r.StaticFS("/image", http.Dir("./image"))

	r.Use(middleware.Cors())

	registerRouter(r)

	go func() {
		err := r.Run(":5500")
		if err != nil {
			fmt.Println(err)
		}
	}()
	
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
