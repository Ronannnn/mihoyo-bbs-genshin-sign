package server

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func NewPrometheusGinExporter(router *gin.Engine) {
	if conf.System.EnablePrometheusMetrics {
		log.Infof("Prometheus metrics http server initialized at %s", conf.System.PrometheusMetricsAddr)
		// New http server for metrics
		metricRouter := gin.New()
		m := ginmetrics.GetMonitor()
		m.SetMetricPath("/metrics")
		m.UseWithoutExposingEndpoint(router)
		m.Expose(metricRouter)
		_ = metricRouter.Run(conf.System.PrometheusMetricsAddr)
	}
}
