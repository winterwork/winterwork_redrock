package controller

import (
	"douban/service"
	"douban/tool"
	"github.com/gin-gonic/gin"
	"strings"
)

// MemberRouter 注册路由
func MemberRouter(r *gin.Engine) {
	engine := r.Group("/member")
	{
		engine.POST("/find", findMember)
		engine.POST("/showMember", showMember)
	}
}

// findMember 根据名字寻找IMDB
func findMember(c *gin.Context) {
	name := c.PostForm("name")
	db := tool.GetDb()
	IMDB, err := service.FindMemberIMDB(db, name)
	tool.CheckErr(err)
	if err != nil {
		if err.Error() == "NoMember" {
			tool.PrintInfo(c, "NoMember!", false)
		}
		return
	}
	IMDBList := ""
	var build strings.Builder
	build.WriteString(IMDBList)
	for _, i := range IMDB {
		if IMDBList != "" {
			build.WriteString("/")
		}
		build.WriteString(i)
		IMDBList = build.String()
	}
	tool.PrintInfo(c, IMDBList, true)
}

// showMember 根据IMDB输出影人信息
func showMember(c *gin.Context) {
	IMDB := c.PostForm("IMDB")
	heading := c.DefaultPostForm("heading", "info")
	db := tool.GetDb()
	u, err := service.ShowMember(db, IMDB, heading)
	tool.CheckErr(err)
	if err != nil {
		tool.PrintInfo(c, "输入了错误的IMDB！", false)
		return
	}
	tool.PrintMember(c, u, heading)
}

