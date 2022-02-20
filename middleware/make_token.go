package middleware

import (
	"douban/tool"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// GenerateToken 签发jwt
func GenerateToken(c *gin.Context, username string) {
	claims := tool.CustomClaim{
		Name: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "kaori.top",                     // 签名颁发者
		},
	}
	token, err := tool.CreateToken(claims)
	tool.CheckErr(err)
	if err != nil {
		return
	}
	fmt.Println(token)
	tool.PrintInfo(c, token, true)
	return
}
