package util

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-gin/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateToken(username, password string) (string, error) {
	var claims Claims
	if password != "" {
		claims = Claims{
			username,
			password,
			jwt.RegisteredClaims{
				// A usual scenario is to set the expiration time relative to the current time
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "test",
				Subject:   "somebody",
				ID:        "1",
				Audience:  []string{"somebody_else"},
			},
		}

	} else {
		claims = Claims{
			username,
			"",
			jwt.RegisteredClaims{
				// A usual scenario is to set the expiration time relative to the current time
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "test",
				Subject:   "somebody",
				ID:        "1",
				Audience:  []string{"somebody_else"},
			},
		}
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	fmt.Printf("%v %v", token, err)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims.Valid {
		fmt.Println("token合法")
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		fmt.Println("传入的字符串甚至连一个token都不是...")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		fmt.Println("token已经过期或者还没有生效")
	} else {
		fmt.Println("token处理异常...")
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Printf("%v %v", claims.Username, claims.RegisteredClaims.Issuer)
			return claims, nil
		}
	}

	return nil, err
}
