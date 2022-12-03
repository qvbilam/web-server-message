package router

import (
	"github.com/gin-gonic/gin"
	"video/api/category"
)

func InitCategoryRouter(Router *gin.RouterGroup) {
	VideoRouter := Router.Group("video/category")
	{
		VideoRouter.GET("", category.List)
		VideoRouter.GET("/:id", category.Detail)
		VideoRouter.POST("", category.Create)
		VideoRouter.PUT("/:id", category.Update)
		VideoRouter.DELETE("/:id", category.Delete)
	}
}
