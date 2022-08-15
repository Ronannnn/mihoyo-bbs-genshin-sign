package main

import (
	"mihoyo-bbs-genshin-sign/config"
	"mihoyo-bbs-genshin-sign/server"
)

var log = config.Logger

func main() {
	log.Info("cron started")
	if c, err := server.NewCronTask(); err != nil {
		log.Error(err)
	} else {
		c.Run()
	}
}
