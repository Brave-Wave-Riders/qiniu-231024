package svc

import (
	"TokTik/app/user/cmd/rpc/internal/config"
	"TokTik/app/user/model/fans"
	"TokTik/app/user/model/users"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"log"
)

type ServiceContext struct {
	Config    config.Config
	UserModel users.UsersModel
	FansModel fans.FansModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gormc.ConnectMysql(c.Mysql)
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:    c,
		UserModel: users.NewUsersModel(db, c.Cache),
		FansModel: fans.NewFansModel(db, c.Cache),
	}
}
