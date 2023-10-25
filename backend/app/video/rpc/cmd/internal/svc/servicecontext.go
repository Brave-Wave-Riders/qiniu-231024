package svc

import (
	"TokTik/app/video/rpc/cmd/internal/config"
	"TokTik/app/video/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Model  *xorm.Engine // 手动代码
	Redis  redis.RedisConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.Init("root:admin123@/qiniuyun?charset=utf8"),
		Redis:  c.Redis.RedisConf,
	}
}
