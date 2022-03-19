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
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "kaori.top",
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
