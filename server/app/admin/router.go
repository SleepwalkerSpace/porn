package admin

import (
	"porn/server/app/admin/handle"
	"porn/server/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(logger))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.Any("/")
			}
			user := v1.Group("/user")
			{
				user.Any("/")
			}
			moive := v1.Group("/moive")
			{
				moive.Any("/")
			}
			actor := v1.Group("/actor")
			{
				actor.Any("/")
			}
			tag := v1.Group("/tag")
			{
				tag.GET("", handle.TagList)
				tag.POST("", handle.TagCreate)
				tag.PUT("", handle.TagUpdate)
				tag.DELETE("", handle.TagDelete)
			}
		}
	}

	return router
}
