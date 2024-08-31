package data

import (
	"github.com/asynccnu/ccnu-service/internal/biz"
	"github.com/asynccnu/ccnu-service/internal/conf"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	logger3 "log"
	"os"
	"time"
)

type Data struct {
	DB *gorm.DB
}

func NewData(c *conf.Data, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return &Data{DB: db}, cleanup, nil
}

// NewDB 连接mysql数据库
func NewDB(c *conf.Data) *gorm.DB {
	newlogger := logger2.New(
		logger3.New(os.Stdout, "\r\n", logger3.LstdFlags),
		logger2.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger2.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open(c.DatabaseSource), &gorm.Config{Logger: newlogger})
	if err != nil {
		panic("connect mysql failed")
	}
	if err := db.AutoMigrate(&biz.User{}); err != nil {
		panic(err)
	}
	return db
}

var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewDB)
