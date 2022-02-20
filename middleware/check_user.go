package middleware

import (
	"douban/tool"
	"github.com/gin-gonic/gin"
)

// Check 检查登录状态
func Check(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		tool.PrintInfo(c, "你还没有登录！", false)
		c.Abort()
		return
	}
	claims, err := tool.TokenParser(token)
	if err != nil {
		if err.Error() == "tokenPass" {
			tool.PrintInfo(c, "你的登录已过期，请重新登录！", false)
			c.Abort()
			return
		}
		tool.CheckErr(err)
		c.Abort()
		return
	}
	c.Set("claims", claims)
}
