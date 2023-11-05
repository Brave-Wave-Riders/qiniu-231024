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

	"log"
)

// Token
//
//	@Description: get
func Token(filename string) string {
	cfg := settings.Conf.QiNiuCloud
	bucket := cfg.Bucket
	fops := cfg.Fops
	saveMp4Entry := base64.URLEncoding.EncodeToString([]byte(bucket + ":" + filename))
	fops += "|saveas/" + saveMp4Entry
	putPolicy := storage.PutPolicy{
		Scope:         bucket,
		PersistentOps: fops,
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
	Name    string `json:"name"`
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
		ctx.JSON(400, gin.H{
			"code":   1,
			"errmsg": "invalid param",
		})
		return
	}
	db := mysql.Get()
	tx := db.Create(info)
	if tx.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   1,
			"errmsg": "database error",
		})
		zap.L().Error(fmt.Sprintf("database error: %s", tx.Error.Error()))
		return
	}
	token := Token(info.Name)
	ctx.JSON(200, gin.H{
		"code":  0,
		"token": token,
		"base":  settings.Conf.QiNiuCloud.Base,
	})
}

// List
//
//	@Description:
//	@param ctx
func List(ctx *gin.Context) {
	pages := ctx.DefaultQuery("pg", "0")
	pagesInt, _ := strconv.Atoi(pages)

	perNums := ctx.DefaultQuery("pgsize", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	if pagesInt <= 0 {
		pagesInt = 1
	}
	if perNumsInt > 35 {
		perNumsInt = 35
	} else if perNumsInt < 15 {
		perNumsInt = 15
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
		ctx.JSON(200, res)
		return
	}
	err := db.Limit(perNumsInt).Offset(offset).Find(&videos).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":   1,
			"errmsg": "database error",
		})
		zap.L().Error(err.Error())
		return
	}
	for _, v := range videos {
		user := &mysql.User{}
		user.ID = int32(v.Uid)
		if err := db.Find(user).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":   1,
				"errmsg": "database error",
			})
			zap.L().Error(err.Error())
			return
		}
		videoList.videos = append(videoList.videos, &VideoItem{Name: v.Name, ImgUrl: v.ImgUrl, Author: user.NickName, UserImg: user.Url})
	}

	res["data"] = &videoList.videos
	res["total"] = count
	rsp, err := json.Marshal(res)
	log.Println(string(rsp))
	//ctx.Data(200, "application/json", rsp)
	ctx.JSON(200, res)
}
