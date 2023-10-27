package svc

import (
	"TokTik/app/user/cmd/api/internal/config"
	"TokTik/app/user/cmd/rpc/userrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpcClient userrpc.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpcClient: userrpc.NewUserrpc(zrpc.MustNewClient(c.UserRpcConfig)),
	}
}
