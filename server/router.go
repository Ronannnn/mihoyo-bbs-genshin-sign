package server

import (
	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors())
	return router
}
