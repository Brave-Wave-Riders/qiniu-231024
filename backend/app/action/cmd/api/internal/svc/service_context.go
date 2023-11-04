package svc

import (
	"TokTik/app/action/cmd/api/internal/config"
	"TokTik/app/action/cmd/rpc/actionrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config          config.Config
	ActionRpcClient actionrpc.Actionrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ActionRpcClient: actionrpc.NewActionrpc(zrpc.MustNewClient(c.ActionRpcClient)),
	}
}
