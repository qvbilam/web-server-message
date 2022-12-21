package router

import (
	"github.com/gin-gonic/gin"
	"message/api"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message/")
	{
		MessageRouter.POST("private/publish", api.PrivateSend)

		MessageRouter.POST("group/publish", api.GroupSend)
	}
}
