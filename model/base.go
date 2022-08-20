package model

import (
	"errors"
	"gorm.io/gorm"
	"mihoyo-bbs-genshin-sign/config"
	"time"
)

var log = config.Logger

var (
	ErrCreatedId = errors.New("id in created entity must be 0")
	ErrUpdatedId = errors.New("id in updated or deleted entity must be greater than 0")
	ErrModified  = errors.New("data modified by others, please refresh the page")
)

var Entities = map[string]interface{}{
	"signItems": &SignItem{},
}

type Base struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement:true"`
	CreatedBy int       `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy int       `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func migrate(db *gorm.DB) {
	for tableName, entity := range Entities {
		if err := db.AutoMigrate(entity); err != nil {
			log.Warnf("Fail to migrate table [%s], try again, %s", tableName, err)
			time.Sleep(time.Second)
			if err := db.AutoMigrate(entity); err != nil {
				panic(err)
			}
		}
	}
}

func init() {
	migrate(config.Db)
}
