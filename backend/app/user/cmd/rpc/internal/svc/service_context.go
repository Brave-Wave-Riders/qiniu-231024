package svc

import (
	"TokTik/app/user/cmd/rpc/internal/config"
	"TokTik/app/user/model"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(db, c.Cache),
	}
}
