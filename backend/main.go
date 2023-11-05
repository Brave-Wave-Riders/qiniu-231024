package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"qiniu_video/dao/mysql"
	"qiniu_video/logger"
	"qiniu_video/routes"
	"qiniu_video/settings"
	"syscall"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed, error desc: %v\n", err)
		return
	}
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, error desc: %v\n", err)
		return
	}
	zap.L().Debug("test logger after init...")

	if _, err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, error desc: %v\n", err)
		return
	}

	//if err := redis.Init(settings.Conf.RedisConfig); err != nil {
	//	fmt.Printf("init redis failed, error desc: %v\n", err)
	//	return
	//}
	//defer redis.Close()

	r := routes.SetUp(settings.Conf.Mode)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		zap.L().Debug(fmt.Sprintf("server start at port: %d", settings.Conf.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: ", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("server shutdown...")

	if err := srv.Shutdown(context.Background()); err != nil {
		zap.L().Fatal(fmt.Sprintf("srv.Shutdown err: %v\n", err))
	}

	zap.L().Info("server exit...")
}
