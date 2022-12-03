package router

import (
	"github.com/gin-gonic/gin"
	"video/api/drama"
	"video/api/video"
)

func InitDramaRouter(Router *gin.RouterGroup) {
	DramaRouter := Router.Group("drama")
	{
		DramaRouter.GET("", drama.List)
		DramaRouter.GET("/:id", drama.Detail)
		DramaRouter.POST("", drama.Create)
		DramaRouter.PUT("/:id", drama.Update)
		DramaRouter.DELETE("/:id", drama.Delete)

		// 剧集
		DramaRouter.GET("/:id/video", video.Detail)
		DramaRouter.POST("/:id/video", video.Create)
		DramaRouter.PUT("/:id/video/:videoId", video.Update)
		DramaRouter.DELETE("/:id/video/:videoId", video.Delete)
	}
}
