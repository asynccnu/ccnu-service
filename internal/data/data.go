package data

import (
	"ccnu-service/internal/conf"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	DB *gorm.DB
}

func NewData(c *conf.Data) (*Data, func(), error) {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return &Data{DB: db}, cleanup, nil
}

var ProviderSet = wire.NewSet(NewData)
