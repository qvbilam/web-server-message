package router

import (
	"github.com/gin-gonic/gin"
	"message/api"
)

func InitPrivateRouter(Router *gin.RouterGroup) {
	PrivateRouter := Router.Group("message/private")
	{
		PrivateRouter.POST("publish", api.PrivateSend)
	}
}
