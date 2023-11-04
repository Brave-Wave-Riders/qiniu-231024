package config

import (
	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql       gormc.Mysql
	Cache       cache.CacheConf
	NacosConfig NacosConfig
}

type NacosConfig struct {
	Host                string
	Port                uint64
	Username            string
	Password            string
	NamespaceId         string
	TimeoutMs           uint64
	NotLoadCacheAtStart bool
	LogLevel            string
}
