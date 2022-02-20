package controller

import (
	"douban/tool"
	"github.com/gin-gonic/gin"
)

// AdminRouter 注册路由
func AdminRouter(engine *gin.Engine) {
	engine.POST("/userDelete", AdminMiddleWare, userDelete)
}

// AdminMiddleWare 管理员权限确认
func AdminMiddleWare(c *gin.Context) {
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	if username == "kaori" {
		return
	}
	tool.PrintInfo(c, "你不是管理员！", false)
	c.Abort()
	return
}

// userDelete 删除用户
func userDelete(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	sqlStr := "delete from user where username = ?"
	_, err := db.Exec(sqlStr, username)
	tool.CheckErr(err)
	tool.PrintInfo(c, "注销成功！", true)
}

