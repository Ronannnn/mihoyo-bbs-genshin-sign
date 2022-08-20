package config

import (
	"gorm.io/driver/sqlite"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db *gorm.DB
)

type DbConf struct {
	DbFilename string `mapstructure:"dbFilename"`
}

func initDb(dbConf DbConf) (err error) {
	if Db, err = gorm.Open(sqlite.Open(dbConf.DbFilename), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		Logger.Error("Fail to connect to sqlite.", zap.Error(err))
		panic(err)
	}
	// set further Conf
	sqlDB, err := Db.DB()
	if err != nil {
		Logger.Error("Fail to get sql db.", zap.Error(err))
		panic(err)
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(100)

	Logger.Info("Database initialized")
	return
}
