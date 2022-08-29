package server

import (
	"github.com/gin-gonic/gin"
	"mihoyo-bbs-genshin-sign/internal/config"
)

var log = config.Logger
var conf = config.Conf

func StartNewServer() {
	gin.SetMode(gin.ReleaseMode)
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Panicking... %v", err)
		}
	}()

	appRouter, metricsRouter := newRouter()
	if conf.System.EnablePrometheusMetrics {
		log.Infof("Launching metrics server at %s", conf.System.PrometheusMetricsAddr)
		go func() {
			_ = metricsRouter.Run(conf.System.PrometheusMetricsAddr)
		}()
	}

	log.Infof("Launching web server at %s", conf.System.Addr)
	if err := appRouter.Run(conf.System.Addr); err != nil {
		panic(err)
	}
}
