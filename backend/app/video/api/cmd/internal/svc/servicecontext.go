package svc

import (
	"TokTik/app/video/api/cmd/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
