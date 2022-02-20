package tool

import (
	"douban/Struct"
	"github.com/gin-gonic/gin"
)

// PrintInfo 输出字符串
func PrintInfo(c *gin.Context, str string, status bool) {
	c.JSON(200, gin.H{
		"status": status,
		"info":   str,
	})
}

// PrintMsg 输出评论
func PrintMsg(c *gin.Context, u []Struct.Info) {
	c.JSON(200, gin.H{
		"information": u,
	})
}

// PrintMovie 输出电影信息
func PrintMovie(c *gin.Context, u interface{}, _type string) {
	switch _type {
	case "info":
		if len(u.([]Struct.MovieInfo)) == 0 {
			c.JSON(200, gin.H{
				"info":   "NoMovie",
				"status": false,
			})
		}
		c.JSON(200, gin.H{
			"information": u,
		})
	case "view":
		if len(u.([]Struct.MovieView)) == 0 {
			c.JSON(200, gin.H{
				"info":   "NoMovie",
				"status": false,
			})
		}
		c.JSON(200, gin.H{
			"information": u,
		})
	case "member":
		if len(u.([]Struct.MovieMember)) == 0 {
			c.JSON(200, gin.H{
				"info":   "NoMovie",
				"status": false,
			})
		}
		c.JSON(200, gin.H{
			"information": u,
		})
	case "liker":
		if len(u.([]Struct.MovieLiker)) == 0 {
			c.JSON(200, gin.H{
				"info":   "NoMovie",
				"status": false,
			})
		}
		c.JSON(200, gin.H{
			"information": u,
		})
	}
}

func PrintMember(c *gin.Context, u interface{}, _type string) {
	switch _type {
	case "info":
		c.JSON(200, gin.H{
			"information": u,
		})
	case "view":
		c.JSON(200, gin.H{
			"information": u,
		})
	}
}

func PrintUser(c *gin.Context, u Struct.UserDetail) {
	c.JSON(200, gin.H{
		"information": u,
	})
}
