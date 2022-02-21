package main

import (
	"douban/controller"
	"douban/middleware"
	"douban/tool"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	tool.OpenDb()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://121.41.120.238:8080","http://121.41.120.238:5500"}
	r.Use(cors.New(config))
	
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
