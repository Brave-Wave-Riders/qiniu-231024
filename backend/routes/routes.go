package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiniu_video/logger"
	"qiniu_video/middlewares"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	videoGroup := r.Group("/video")
	videoV1 := videoGroup.Group("/v1")
	{
		videoV1.POST("/upload", middlewares.JwtAuth(), Upload)
		videoV1.GET("/list/:type", List)
		videoV1.GET("/list/my", middlewares.JwtAuth(), MyList)
	}
	userGroup := r.Group("/user")
	userV1 := userGroup.Group("/v1")
	{
		userV1.POST("/login", Login)
		userV1.POST("/register", Register)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
