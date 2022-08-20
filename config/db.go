package config

import (
	"database/sql"
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
	if err = CloseDb(Db); err != nil {
		return
	}
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

func CloseDb(db *gorm.DB) (err error) {
	if db == nil {
		return
	}
	var sqlDb *sql.DB
	if sqlDb, err = db.DB(); err != nil {
		return
	}
	return sqlDb.Close()
}
