package tool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("kaori.top"),
	}
}

type CustomClaim struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// CreateToken 制造token
func CreateToken(chaim CustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, chaim)
	return token.SignedString(NewJWT().SigningKey)
}

// TokenParser 解密jwt
func TokenParser(tokenString string) (*CustomClaim, error) {
	j := NewJWT()
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	CheckErr(err)
	if err != nil {
		if err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
			return nil, errors.New("tokenPass")
		}
		return nil, err
	}
	chaim := token.Claims.(*CustomClaim)
	if token.Valid {
		return chaim, nil
	}
	return nil, err
}
