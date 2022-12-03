package router

import (
	"github.com/gin-gonic/gin"
	"video/api/region"
)

func InitRegionRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video/region")
	{
		VideoRouter.GET("", region.List)
		VideoRouter.GET("/:id", region.Detail)
		VideoRouter.POST("", region.Create)
		VideoRouter.PUT("/:id", region.Update)
		VideoRouter.DELETE("/:id", region.Delete)
	}
}
