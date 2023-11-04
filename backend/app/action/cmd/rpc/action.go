package main

import (
	"TokTik/app/action/cmd/rpc/internal/config"
	"TokTik/app/action/cmd/rpc/internal/server"
	"TokTik/app/action/cmd/rpc/internal/svc"
	"TokTik/app/action/cmd/rpc/pb"
	"flag"
	"github.com/Guo-Chenxu/go-zero-nacos/configcenter"
	"github.com/Guo-Chenxu/go-zero-nacos/register"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var bootstrapConfigFile = flag.String("f", "etc/bootstrap.yaml", "the config.go file")

func main() {
	//解析bootstrap并获取业务配置
	flag.Parse()
	var bootstrapConfig configcenter.BootstrapConfig
	var c config.Config

	configcenter.GetConfig(bootstrapConfigFile, &bootstrapConfig, &c)

	// 进行业务配置
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterActionrpcServer(grpcServer, server.NewActionrpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 注册nacos
	register.Register(c)

	s.Start()
}
