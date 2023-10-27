package main

import (
	"TokTik/app/user/cmd/api/internal/config"
	"TokTik/app/user/cmd/api/internal/handler"
	"TokTik/app/user/cmd/api/internal/svc"
	"flag"
	"fmt"
	"github.com/Guo-Chenxu/go-zero-nacos/configcenter"
	"github.com/zeromicro/go-zero/rest"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/nacos"
)

var bootstrapConfigFile = flag.String("f", "etc/bootstrap.yaml", "the config.go file")

func main() {
	//解析bootstrap config.go
	flag.Parse()
	var bootstrapConfig configcenter.BootstrapConfig
	var c config.Config
	configcenter.GetConfig(bootstrapConfigFile, &bootstrapConfig, &c)

	// 配置业务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
