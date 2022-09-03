package admin

import (
	"porn/server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(logger))
	router.GET("/", func(ctx *gin.Context) {

		ctx.String(200, "hi")
	})
	return router
}
