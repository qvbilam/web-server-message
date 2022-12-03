package router

import (
	"github.com/gin-gonic/gin"
	"video/api/episode"
)

func InitEpisodesRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video/episodes")
	{
		VideoRouter.GET("", episode.List)
		VideoRouter.GET("/:id", episode.Detail)
		VideoRouter.POST("", episode.Create)
		VideoRouter.PUT("/:id", episode.Update)
		VideoRouter.DELETE("/:id", episode.Delete)
	}
}
