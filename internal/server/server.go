package server

import (
	"context"
	"mihoyo-bbs-genshin-sign/internal/config"
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

	router := newRouter()
	server := &http.Server{
		Addr:    conf.System.Addr,
		Handler: router,
	}

	go func() {
		log.Infof("Launching web server[v%s] at %s", config.Version, server.Addr)

		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Info("Web server shutdown completely")
			} else {
				log.Error("Web server closed with exceptions", err)
			}
		}
	}()

	go NewPrometheusGinExporter(router)

	go func() {
		if cron, err := NewCronTask(); err != nil {
			log.Error("Fail to create cron task", err)
		} else {
			cron.Run()
		}
	}()

	<-ctx.Done()
	log.Info("http: shutting down the server")
	err := server.Close()

	if err != nil {
		log.Error("Fail to shutdown the server", err)
	}
}
