package svc

import (
	"TokTik/app/video/api/cmd/internal/config"
	"TokTik/app/video/rpc/cmd/videoclient"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Video          videoclient.Video
	KqPusherClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		Video:          videoclient.NewVideo(zrpc.MustNewClient(c.Video)),
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
