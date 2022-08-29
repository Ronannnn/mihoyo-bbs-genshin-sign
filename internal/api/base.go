package api

import (
	"github.com/google/uuid"
	"mihoyo-bbs-genshin-sign/internal/config"
)

var (
	db  = config.Db
	log = config.Logger
)

var (
	ServerId = uuid.New().String()
)
