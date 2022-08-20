package server

import (
	"context"
	"mihoyo-bbs-genshin-sign/config"
	"net/http"

	"go.uber.org/zap"
)

var log = config.Logger
var conf = config.Conf

func NewServer(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Panicking...", zap.Any("err", err))
		}
	}()

	server := &http.Server{
		Addr:    conf.System.Addr,
		Handler: newRouter(),
	}

	go func() {
		log.Infof("Launching web server[v%s] at %s", config.Version, server.Addr)

		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("Web server shutdown completely")
			} else {
				log.Error("Web server closed with exceptions", zap.Error(err))
			}
		}
	}()

	go func() {
		if cron, err := NewCronTask(); err != nil {
			log.Error("Fail to create cron task", zap.Error(err))
		} else {
			cron.Run()
		}
	}()

	<-ctx.Done()
	log.Info("http: shutting down web server")
	err := server.Close()

	if err != nil {
		log.Error("Fail to shutdown the server", zap.Error(err))
	}
}
