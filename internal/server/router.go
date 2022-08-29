package server

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"mihoyo-bbs-genshin-sign/internal/api"
)

func newRouter() (appRouter *gin.Engine, metricsRouter *gin.Engine) {
	appRouter = gin.New()
	appRouter.Use(cors())

	// must be placed before endpoints
	// or later endpoints will not use this handler
	// see https://github.com/gin-gonic/gin/issues/1224 for more details
	metricsRouter = gin.New()
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.UseWithoutExposingEndpoint(appRouter)
	m.Expose(metricsRouter)

	apiRouter := appRouter.Group("/api")
	apiV1Router := apiRouter.Group("/v1")
	{
		apiV1Router.GET("/sign", api.GetAllSignItems)
		apiV1Router.POST("/sign", api.CreateSignItem)
		apiV1Router.PUT("/sign", api.UpdateSignItem)
		apiV1Router.DELETE("/sign/:id", api.DeleteSignItemById)
	}

	return
}
