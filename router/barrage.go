package router

import (
	"github.com/gin-gonic/gin"
	"video/api/barrage"
)

func InitBarrageRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video/barrage")
	{
		VideoRouter.GET("", barrage.List)
		VideoRouter.GET("/:id", barrage.Detail)
		VideoRouter.POST("", barrage.Create)
		VideoRouter.PUT("/:id", barrage.Update)
		VideoRouter.DELETE("/:id", barrage.Delete)
	}
}
