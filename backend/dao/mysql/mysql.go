package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm/logger"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"qiniu_video/settings"
)

var (
	once sync.Once
	db   *gorm.DB
)

func Init(cfg *settings.MySQLConfig) (*gorm.DB, error) {
	var err error
	once.Do(func() {
		if cfg == nil {
			err = fmt.Errorf("no mysql options, get failed.")
			return
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			return
		}

		var sqlDB *sql.DB
		sqlDB, err = db.DB()

		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Second)

	})
	if db == nil && err != nil {
		return nil, errors.New("get db failed.")
	}
	db.AutoMigrate(&Video{})
	db.AutoMigrate(&User{})
	return db, nil
}

func Get() *gorm.DB {
	return db
}
