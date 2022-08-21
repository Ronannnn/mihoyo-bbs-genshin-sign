package server

import (
	"github.com/robfig/cron/v3"
	"mihoyo-bbs-genshin-sign/config"
	"mihoyo-bbs-genshin-sign/service"
)

func NewCronTask() (c *cron.Cron, err error) {
	c = cron.New()
	// utc every day 00:00 -> utc+8 every day 8:00
	if _, err = c.AddFunc("0 0 * * *", func() {
		if err = service.SignCronTask(config.Db); err != nil {
			log.Error(err)
		} else {
			log.Info("signed")
		}
	}); err != nil {
		return nil, err
	}
	return
}
