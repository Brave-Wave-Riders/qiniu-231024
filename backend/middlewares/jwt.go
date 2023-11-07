package middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"qiniu_video/settings"
	"strings"
	"time"
)

type MyCustomClaims struct {
	Nickname string `json:"username"`
	jwt.StandardClaims
}

func JwtDecode(tokenString string) (*MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.Conf.JwtConfig.Key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("request expired")
		}

		return claims, nil
	}

	return nil, err
}

func JwtEncode(claims MyCustomClaims) (string, error) {
	mySigningKey := []byte(settings.Conf.JwtConfig.Key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", 1)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
				"errmsg": "请登录",
				"code":   1,
			})
			ctx.Abort()
			return
		}
		claims, err := JwtDecode(token)
		if err != nil || claims == nil {
			zap.L().Error(fmt.Sprintf("JwtDecode failed: %v", err))
			ctx.JSON(200, map[string]interface{}{
				"errmsg": "内部错误，token解析失败",
				"code":   2,
			})
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Nickname)
	}
}
