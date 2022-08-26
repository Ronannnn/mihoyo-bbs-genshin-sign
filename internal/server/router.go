package server

import (
	"github.com/gin-gonic/gin"
	"mihoyo-bbs-genshin-sign/internal/api"
)

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors())
	apiRouter := router.Group("/api")
	apiV1Router := apiRouter.Group("/v1")
	{
		apiV1Router.GET("/sign", api.GetAllSignItems)
		apiV1Router.POST("/sign", api.CreateSignItem)
		apiV1Router.PUT("/sign", api.UpdateSignItem)
		apiV1Router.DELETE("/sign/:id", api.DeleteSignItemById)
	}

	return router
}
