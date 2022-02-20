package controller

import (
	"douban/Struct"
	"douban/middleware"
	"douban/service"
	"douban/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

// MessageRouter 注册路由
func MessageRouter(r *gin.Engine) {
	engine := r.Group("/message")
	{
		engine.POST("/sendMsg", middleware.Check, SendMsg)
		engine.POST("/anonymousMsg", middleware.Check, anonymousMsg)
		engine.POST("/msgDelete", middleware.Check, AdminMiddleWare, deleteMsg)
		engine.POST("/msgList", listMsg)
		engine.POST("/msg", getOneMsg)
		engine.POST("/comment", middleware.Check, SendComment)
		engine.POST("/messages", listComment)
		engine.POST("/thumb", middleware.Check, thumb)
		engine.POST("/essay", middleware.Check, sendEssay)
		engine.POST("/pointOnly", middleware.Check, pointMovie)
		engine.POST("/deletePoint", middleware.Check, deletePoint)
		engine.POST("/deleteThumb", middleware.Check, deleteThumb)
		engine.POST("/thumbUser", middleware.Check, thumbUser)
		engine.POST("/deleteUT", middleware.Check, deleteUT)
	}
}

// thumb 点赞
func thumb(c *gin.Context) {
	id := c.PostForm("id")
	db := tool.GetDb()
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	err := service.ThumbAdd(db, id, username)
	if err != nil {
		if err.Error() == "done" {
			tool.PrintInfo(c, "你已经赞过了！", false)
		}
		tool.PrintInfo(c, "找不到该id捏！", false)
		return
	}
	tool.PrintInfo(c, "好力！", true)
}

// SendMsg 发送留言
func SendMsg(c *gin.Context) {
	msg := c.PostForm("msg")
	claims, _ := c.Get("claims")
	movie := c.PostForm("movie")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	id, err := service.SendMsg(db, username, msg, 0, 1, movie, 0)
	tool.CheckErr(err)
	if err != nil {
		tool.PrintInfo(c, "出问题了捏！", false)
	}
	tool.PrintInfo(c, strconv.Itoa(id), true)
	return
}

// anonymousMsg 发送匿名留言
func anonymousMsg(c *gin.Context) {
	msg := c.PostForm("msg")
	movie := c.PostForm("movie")
	username := "anonymousUser"
	db := tool.GetDb()
	id, err := service.SendMsg(db, username, msg, 0, 1, movie, 0)
	tool.CheckErr(err)
	tool.PrintInfo(c, strconv.Itoa(id), true)
	return
}

// getOneMsg 获取一条留言
func getOneMsg(c *gin.Context) {
	id := c.PostForm("id")
	db := tool.GetDb()
	u0, err := service.GetOneMsg(db, id)
	if err != nil {
		if err.Error() == "NoMessage" {
			tool.PrintInfo(c, "无该id对于的留言。", false)
		}
		return
	}
	u := []Struct.Info{u0}
	tool.PrintMsg(c, u)
}

// listMsg 获取全部留言
func listMsg(c *gin.Context) {
	db := tool.GetDb()
	movie := c.PostForm("movie")
	u, err := service.FindAllPMsg(db, 0, movie)
	tool.CheckErr(err)
	if err != nil {
		tool.PrintInfo(c, "No message!", false)
		return
	}
	tool.PrintMsg(c, u)
	return
}

// SendComment 发送评论
func SendComment(c *gin.Context) {
	msg := c.PostForm("msg")
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	pid, err := strconv.Atoi(c.PostForm("toWhich"))
	tool.CheckErr(err)
	if err != nil {
		tool.PrintInfo(c, "你输入的评论地址有误！", false)
		return
	}
	db := tool.GetDb()
	id, err := service.SendMsg(db, username, msg, pid, 1, "comment", 0)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	tool.PrintInfo(c, strconv.Itoa(id), true)
}

// listComment 显示所有评论
func listComment(c *gin.Context) {
	db := tool.GetDb()
	pid := c.PostForm("toWhich")
	pidNum, _ := strconv.Atoi(pid)
	u, err := service.FindAllComments(db, pidNum)
	if err != nil {
		if err.Error() == "NoMessage" {
			tool.PrintInfo(c, "No message!", false)
		}
		return
	}
	tool.PrintMsg(c, u)
	return
}

// deleteMsg 删除留言
func deleteMsg(c *gin.Context) {
	id := c.PostForm("id")
	db := tool.GetDb()
	err := service.DeleteMsg(db, id)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	tool.PrintInfo(c, "删除成功！", true)
}

// sendEssay 发送留言
func sendEssay(c *gin.Context) {
	msg := c.PostForm("essay")
	claims, _ := c.Get("claims")
	movie := c.PostForm("movie")
	point := c.PostForm("point")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	pointNum, _ := strconv.Atoi(point)
	id, err := service.SendMsg(db, username, msg, 0, 2, movie, pointNum)
	tool.CheckErr(err)
	if err != nil {
		tool.PrintInfo(c, "", false)
		return
	}
	tool.PrintInfo(c, strconv.Itoa(id), true)
	return
}

// pointMovie 关注/看过电影
func pointMovie(c *gin.Context) {
	_type := c.PostForm("type")
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	movie := c.PostForm("movie")
	err := service.UserPoint(db, username, _type, movie)
	tool.CheckErr(err)
	if err != nil {
		if err.Error() == "done" {
			tool.PrintInfo(c, "你已经关注/标注看过这部电影了！", false)
		}
		return
	}
	tool.PrintInfo(c, "好力！", true)
}

// deleteThumb 取消评论点赞
func deleteThumb(c *gin.Context) {
	ID := c.PostForm("id")
	IDNum, _ := strconv.Atoi(ID)
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	err := service.DeleteThumb(db, username, IDNum)
	tool.CheckErr(err)
	if err != nil {
		if err.Error() == "NoThumb" {
			tool.PrintInfo(c, "你还没有赞过该条评论！", false)
		}
		return
	}
	tool.PrintInfo(c, "已成功取消点赞！", true)
}

// deletePoint 删除原有的评分或想看/看过
func deletePoint(c *gin.Context) {
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	movie := c.PostForm("movie")
	_type := c.PostForm("type")
	err := service.DeletePoint(db, username, movie, _type)
	tool.CheckErr(err)
	if err != nil {
		if err.Error() == "NoPoint" {
			tool.PrintInfo(c, "你还没有想看/看过", false)
		}
		return
	}
	tool.PrintInfo(c, "好力！", true)
}

// thumbUser 点赞用户
func thumbUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	user := c.PostForm("user")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	err := service.ThumbUser(db, user, username)
	tool.CheckErr(err)
	if err != nil {
		if err.Error() == "done" {
			tool.PrintInfo(c, "你已经赞过了！", false)
		}
		return
	}
	tool.PrintInfo(c, "好力！", true)
}

// deleteUT 删除对用户的点赞
func deleteUT(c *gin.Context) {
	user := c.PostForm("user")
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	err := service.DeleteUT(db, user, username)
	tool.CheckErr(err)
	if err != nil {
		if err.Error() == "NoThumb" {
			tool.PrintInfo(c, "你还没有赞过！", false)
		}
		return
	}
	tool.PrintInfo(c, "好力！", true)
}
