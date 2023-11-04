package svc

import (
	"TokTik/app/action/cmd/rpc/internal/config"
	"TokTik/app/action/model/fans"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	FansModel fans.FansModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		FansModel: fans.NewFansModel(db, c.Cache),
	}
}
