package router

import (
	"github.com/gin-gonic/gin"
	"video/api/video"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video")
	{
		VideoRouter.GET("", video.List)
		VideoRouter.GET("/:id", video.Detail)
		VideoRouter.POST("", video.Create)
		VideoRouter.PUT("/:id", video.Update)
		VideoRouter.DELETE("/:id", video.Delete)
	}
}
