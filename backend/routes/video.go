package routes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"net/http"
	"qiniu_video/dao/mysql"
	"qiniu_video/settings"
	"strconv"
	"strings"

	"log"
)

// Token
//
//	@Description: get
func Token(filename string) string {
	//"persistentOps":        "vframe/jpg/offset/1/w/480/h/360",
	cfg := settings.Conf.QiNiuCloud
	bucket := cfg.Bucket
	fops := cfg.Fops
	saveMp4Entry := base64.URLEncoding.EncodeToString([]byte(bucket + ":" + filename))
	saveJpgEntry := base64.URLEncoding.EncodeToString([]byte(cfg.Bucket + ":" + filename + ".jpg"))
	fops += "|saveas/" + saveMp4Entry
	imgFops := cfg.ImgFops + "|saveas/" + saveJpgEntry
	persistentOps := strings.Join([]string{fops, imgFops}, ";")
	putPolicy := storage.PutPolicy{
		Scope:         bucket,
		PersistentOps: persistentOps,
	}
	putPolicy.Expires = cfg.Expires
	accessKey := cfg.AccessKey
	secretKey := cfg.SecretKey
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	log.Println(upToken)
	return upToken
}

type VideoItem struct {
	ImgUrl  string `json:"url"`
	Name    string `json:"desc"`
	Author  string `json:"author"`
	UserImg string `json:"icon"`
}

type VideoList struct {
	videos []*VideoItem `json:"videos"`
}

// Upload
//
//	@Description:
//	@param ctx
func Upload(ctx *gin.Context) {
	info := &mysql.Video{}
	if err := ctx.ShouldBind(info); err != nil {
		ctx.JSON(400, &Response{
			Code:   1,
			ErrMsg: "invalid param",
		})
		return
	}
	db := mysql.Get()
	tx := db.Create(info)
	if tx.Error != nil {
		ctx.JSON(http.StatusOK, &Response{
			Code:   1,
			ErrMsg: "database error",
		})
		zap.L().Error(fmt.Sprintf("database error: %s", tx.Error.Error()))
		return
	}
	token := Token(info.Name)
	data := make(map[string]interface{})
	data["token"] = token
	data["base"] = settings.Conf.QiNiuCloud.Base

	ctx.JSON(200, &Response{
		Code: 0,
		Data: data,
	})
}

// List
//
//	@Description:
//	@param ctx
func List(ctx *gin.Context) {
	pages := ctx.DefaultQuery("page", "0")
	pagesInt, _ := strconv.Atoi(pages)

	perNums := ctx.DefaultQuery("size", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	if pagesInt <= 0 {
		pagesInt = 1
	}
	if perNumsInt > 35 {
		perNumsInt = 35
	} else if perNumsInt < 10 {
		perNumsInt = 10
	}
	offset := (pagesInt - 1) * perNumsInt
	vtypeStr := ctx.Params.ByName("type")
	vtype := 0
	if vtypeStr != "" {
		vtype, _ = strconv.Atoi(vtypeStr)
	}

	videos := []*mysql.Video{}
	videoList := &VideoList{}
	db := mysql.Get()
	if vtype != 0 {
		db = db.Where("video_type = ?", vtype)
	}
	var count int64
	res := make(map[string]interface{})
	db.Model(&mysql.Video{}).Count(&count)
	if count == 0 {
		res["total"] = 0
		ctx.JSON(200, &Response{
			Code: 0,
			Data: res,
		})
		return
	}
	err := db.Limit(perNumsInt).Offset(offset).Find(&videos).Error
	if err != nil {
		ctx.JSON(http.StatusOK, &Response{
			Code:   1,
			ErrMsg: "database error",
		})
		zap.L().Error(err.Error())
		return
	}
	for _, v := range videos {
		user := mysql.User{}
		//user.ID = int32(v.Uid)
		if err := mysql.Get().First(&user, v.Uid).Error; err != nil {
			ctx.JSON(http.StatusOK, &Response{
				Code:   1,
				ErrMsg: "database error",
			})
			zap.L().Error(err.Error())
			return
		}
		videoList.videos = append(videoList.videos, &VideoItem{Name: v.Name, ImgUrl: v.ImgUrl, Author: user.NickName, UserImg: user.Url})
	}

	res["data"] = &videoList.videos
	res["total"] = count
	res["base"] = settings.Conf.QiNiuCloud.Base
	rsp, err := json.Marshal(res)
	log.Println(string(rsp))
	//ctx.Data(200, "application/json", rsp)
	ctx.JSON(200, &Response{
		Code: 0,
		Data: res,
	})
}

func MyList(ctx *gin.Context) {
	pages := ctx.DefaultQuery("page", "0")
	pagesInt, _ := strconv.Atoi(pages)

	perNums := ctx.DefaultQuery("size", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	if pagesInt <= 0 {
		pagesInt = 1
	}
	if perNumsInt > 35 {
		perNumsInt = 35
	} else if perNumsInt < 10 {
		perNumsInt = 10
	}
	offset := (pagesInt - 1) * perNumsInt

	videos := []*mysql.Video{}
	videoList := &VideoList{}
	db := mysql.Get()
	// 我的 创作
	v, ok := ctx.Get("username")
	if !ok {
		ctx.JSON(http.StatusBadRequest, &Response{
			Code:   1,
			ErrMsg: "用户未登录",
		})
		//zap.L().Error(err.Error())
		return
	}
	uname := v.(string)
	udb := mysql.Get()
	user := mysql.User{}
	if err := udb.Where("nick_name = ?", uname).First(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Code:   1,
			ErrMsg: "用户不存在",
		})
		//zap.L().Error(err.Error())
		return
	}
	db = db.Where("uid = ?", user.ID)

	var count int64
	res := make(map[string]interface{})
	db.Model(&mysql.Video{}).Count(&count)
	if count == 0 {
		res["total"] = 0
		ctx.JSON(200, &Response{
			Code: 0,
			Data: res,
		})
		return
	}
	err := db.Limit(perNumsInt).Offset(offset).Find(&videos).Error
	if err != nil {
		ctx.JSON(http.StatusOK, &Response{
			Code:   1,
			ErrMsg: "database error",
		})
		zap.L().Error(err.Error())
		return
	}
	for _, v := range videos {
		user := mysql.User{}
		//user.ID = int32(v.Uid)
		if err := mysql.Get().First(&user, v.Uid).Error; err != nil {
			ctx.JSON(http.StatusOK, &Response{
				Code:   1,
				ErrMsg: "database error",
			})
			zap.L().Error(err.Error())
			return
		}
		videoList.videos = append(videoList.videos, &VideoItem{Name: v.Name, ImgUrl: v.ImgUrl, Author: user.NickName, UserImg: user.Url})
	}

	res["data"] = &videoList.videos
	res["total"] = count
	res["base"] = settings.Conf.QiNiuCloud.Base
	rsp, err := json.Marshal(res)
	log.Println(string(rsp))
	//ctx.Data(200, "application/json", rsp)
	ctx.JSON(200, &Response{
		Code: 0,
		Data: res,
	})
}
