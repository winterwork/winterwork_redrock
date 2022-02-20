package controller

import (
	"douban/service"
	"douban/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

// MovieRouter 注册路由
func MovieRouter(r *gin.Engine) {
	engine := r.Group("/movie")
	{
		engine.POST("/find", findMoviesInfo)
		engine.POST("/getRand", getRandMovie)
		engine.POST("/findByIMDB", findByIMDB)
		engine.POST("/findByMember", findByMember)
	}
}

// findMoviesInfo 查找指定的电影Info
func findMoviesInfo(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	heading := c.DefaultPostForm("heading", "info")
	db := tool.GetDb()
	u, err := service.FindMovies(db, name, heading)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	tool.PrintMovie(c, u, heading)
}

// getRandMovie 查找随机电影
func getRandMovie(c *gin.Context) {
	_type := c.DefaultPostForm("type", "")
	numStr := c.DefaultPostForm("num", "8")
	heading := c.DefaultPostForm("heading", "info")
	num, _ := strconv.Atoi(numStr)
	db := tool.GetDb()
	Movie, err := service.GetRandMovie(db, _type, num, heading)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	tool.PrintMovie(c, Movie, heading)
}

// findByIMDB 根据IMDB查找电影
func findByIMDB(c *gin.Context) {
	name := c.PostForm("IMDB")
	heading := c.DefaultPostForm("heading", "info")
	db := tool.GetDb()
	u, err := service.FindByIMDB(db, name, heading)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	tool.PrintMovie(c, u, heading)
}

// findByMember 根据影人查找IMDB
func findByMember(c *gin.Context) {
	IMDB := c.PostForm("MemberIMDB")
	db := tool.GetDb()
	u, err := service.FindByMIMDB(db, IMDB)
	if err != nil {
		if err.Error() == "NoRows" {
			tool.PrintInfo(c, "没有符合该影人IMDB的电影", false)
		}
	}
	tool.PrintInfo(c, u, true)
}
