package controller

import (
	"douban/dao"
	"douban/middleware"
	"douban/service"
	"douban/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserRouter 注册路由
func UserRouter(r *gin.Engine) {
	engine1 := r.Group("/user")
	{
		engine1.POST("/login", login)
		engine1.POST("/register", register)
		engine1.POST("/change", changePassword)
		engine1.POST("/showSecret", showSecret)
		engine1.POST("/addSecret", middleware.Check, addSecret)
		engine1.POST("/check", middleware.Check, done)
		engine1.POST("/getDetail", getDT)
		engine1.POST("/getID", getID)
	}
	engine2 := r.Group("/homepage")
	{
		engine3 := engine2.Group("/introduce")
		{
			engine3.POST("/update", middleware.Check, updateHP)
			engine3.POST("/get", getHP)
		}
	}
}

// addSecret 添加密保
func addSecret(c *gin.Context) {
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	db := tool.GetDb()
	u, _ := service.QueryRowDemo(db, username)
	if u.Question == "" || u.Answer == "" {
		question := c.PostForm("question")
		answer := c.PostForm("answer")
		err := service.AddSecret(db, question, answer, username)
		tool.CheckErr(err)
		tool.PrintInfo(c, "成功设置！", true)
		return
	}
	tool.PrintInfo(c, "你已经设置过密保！", false)
}

// login 登录
func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("LoginUserInfo:", username, password)
	db := tool.GetDb()
	flag := service.QueryRowDemoPassword(db, username, password)
	if flag {
		middleware.GenerateToken(c, username)
		return
	}
	tool.PrintInfo(c, "账号或密码错误！", false)
}

// register 注册
func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	secretQuestion := c.DefaultPostForm("question", "")
	secretAnswer := c.DefaultPostForm("answer", "")
	if username == "" {
		tool.PrintInfo(c, "用户名不能为空！", false)
		return
	}
	if len(username) > 20 {
		tool.PrintInfo(c, "用户名长度超出限制！", false)
		return
	}
	if len(password) > 16 {
		tool.PrintInfo(c, "密码长度超出限制！", false)
		return
	} else if len(password) < 8 {
		tool.PrintInfo(c, "密码太短！", false)
		return
	}
	db := tool.GetDb()
	_, flag := service.QueryRowDemo(db, username)
	if flag {
		tool.PrintInfo(c, "该账号已经被注册!", false)
		return
	}
	flag2 := service.RegisterUser(db, username, password, secretAnswer, secretQuestion)
	if !flag2 {
		tool.PrintInfo(c, "注册失败！", false)
		return
	}
	tool.PrintInfo(c, "注册成功！", true)
}

// changePassword 修改密码
func changePassword(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	u, flag := service.QueryRowDemo(db, username)
	if flag {
		answer := c.PostForm("answer")
		password := c.PostForm("newPassword")
		if answer == u.Answer {
			err := service.ChangePassword(db, password, username)
			tool.CheckErr(err)
			tool.PrintInfo(c, "修改完成！", true)
		} else {
			tool.PrintInfo(c, "密保答案错误！", false)
			return
		}
	} else {
		password2 := c.PostForm("newPassword")
		err := service.ChangePassword(db, password2, username)
		tool.CheckErr(err)
		tool.PrintInfo(c, "你还没有设置密保，修改完成!", true)
	}
}

// showSecret 展示密保
func showSecret(c *gin.Context) {
	username := c.PostForm("username")
	db := tool.GetDb()
	u, flag := service.QueryRowDemo(db, username)
	if !flag {
		tool.PrintInfo(c, "查无此号", false)
		return
	}
	tool.PrintInfo(c, u.Question, true)
}

// done 检查登录状态
func done(c *gin.Context) {
	u := c.MustGet("claims")
	tool.PrintInfo(c, u.(*tool.CustomClaim).Name, true)
}

// updateHP 更新个人简介
func updateHP(c *gin.Context) {
	claims, _ := c.Get("claims")
	username := claims.(*tool.CustomClaim).Name
	message := c.PostForm("message")
	db := tool.GetDb()
	err := service.UpdateHP(db, username, message)
	tool.CheckErr(err)
	if err != nil {
		tool.PrintInfo(c, "失败", false)
		return
	}
	tool.PrintInfo(c, "上传成功", true)
}

// getHP 获取个人简介
func getHP(c *gin.Context) {
	IDStr := c.PostForm("id")
	db := tool.GetDb()
	ID, _ := strconv.Atoi(IDStr)
	message, err := service.GetHP(db, ID)
	if err != nil {
		if err.Error() == "NoUser" {
			tool.PrintInfo(c, "没有这个用户", false)
		}
		return
	}
	tool.PrintInfo(c, message, true)
}

// getDT 获取用户细节
func getDT(c *gin.Context) {
	id := c.PostForm("id")
	db := tool.GetDb()
	idNum, _ := strconv.Atoi(id)
	u, err := service.GetDT(db, idNum)
	if err != nil {
		tool.PrintInfo(c, "NoUser", false)
		return
	}
	tool.PrintUser(c, u)
}

// getID 根据用户名查询用户ID
func getID(c *gin.Context) {
	name := c.PostForm("username")
	db := tool.GetDb()
	id, err := dao.GetId(db, name)
	if err != nil {
		tool.PrintInfo(c, "NoUser", false)
		return
	}
	tool.PrintInfo(c, strconv.Itoa(id), true)
}
