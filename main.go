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
//e2299bbc41e5bfb23e2fa5c5a13988d05863a4be
func main() {
	tool.OpenDb()

	r := gin.Default()

	r.Use(middleware.Cors())

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	r.StaticFS("/image", http.Dir("./image"))

	registerRouter(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
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
