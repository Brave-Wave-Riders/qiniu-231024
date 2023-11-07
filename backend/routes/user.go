package routes

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"qiniu_video/dao/mysql"
	"qiniu_video/middlewares"
	"qiniu_video/settings"
	"time"
)

type UserInput struct {
	NickName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserRegistInput struct {
	NickName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}

func Login(ctx *gin.Context) {
	input := UserInput{}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "invalid param",
		})
		return
	}
	user := &mysql.User{}
	user.NickName = input.NickName
	db := mysql.Get()
	err := db.Where("nick_name = ?", user.NickName).Find(user).Error
	if err != nil {
		zap.L().Error(fmt.Sprintf("get user failed: %s", err.Error()))
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "用户未找到",
		})
		return
	}

	if user.Password != input.Password {
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "用户密码错误",
		})
		return
	}
	claims := middlewares.MyCustomClaims{
		user.NickName,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //签名的生效时间
			ExpiresAt: time.Now().Local().Add(time.Duration(settings.Conf.JwtConfig.Timeout) * time.Hour).Unix(),
			Issuer:    settings.Conf.JwtConfig.Issuer,
		},
	}
	token, err := middlewares.JwtEncode(claims)
	if err != nil {
		zap.L().Error(fmt.Sprintf("JwtEncode failed: ", err.Error()))
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "token生成失败",
		})
		return
	}
	res := make(map[string]interface{})
	res["token"] = token
	ctx.JSON(200, &Response{
		Code: 0,
		Data: res,
	})
}

func Register(ctx *gin.Context) {
	input := UserRegistInput{}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "invalid param",
		})
		return
	}
	user := &mysql.User{
		NickName: input.NickName,
		Password: input.Password,
		Mobile:   input.Mobile,
	}

	db := mysql.Get()
	err := db.Create(user).Error
	if err != nil {
		zap.L().Error(fmt.Sprintf("create user failed: %s", err.Error()))
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "用户已存在",
		})
		return
	}

	claims := middlewares.MyCustomClaims{
		user.NickName,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //签名的生效时间
			ExpiresAt: time.Now().Local().Add(time.Duration(settings.Conf.JwtConfig.Timeout) * time.Hour).Unix(),
			Issuer:    settings.Conf.JwtConfig.Issuer,
		},
	}
	token, err := middlewares.JwtEncode(claims)
	if err != nil {
		zap.L().Error(fmt.Sprintf("JwtEncode failed: ", err.Error()))
		ctx.JSON(200, &Response{
			Code:   1,
			ErrMsg: "token生成失败",
		})
		return
	}
	res := make(map[string]interface{})
	res["token"] = token
	ctx.JSON(200, &Response{
		Code: 0,
		Data: res,
	})
}
