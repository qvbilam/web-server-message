package router

import (
	"github.com/gin-gonic/gin"
	"message/api/broadcast"
	"message/api/group"
	"message/api/private"
	"message/middleware"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	MessageRouter := Router.Group("message/").Use(middleware.Cors()).Use(middleware.Auth())
	{
		MessageRouter.GET("private/:id", private.Message)
		MessageRouter.POST("private/publish", private.Send)
		MessageRouter.POST("private/publish/txt", private.SendText)
		MessageRouter.POST("private/publish/img", private.SendImage)
		MessageRouter.POST("private/read", private.Read)
		MessageRouter.POST("private/rollback", private.Rollback)

		MessageRouter.GET("group/:id", group.Message)
		MessageRouter.POST("group/publish", group.Send)
		MessageRouter.POST("group/publish/txt", group.SendText)
		MessageRouter.POST("group/publish/img", group.SendImage)

		MessageRouter.POST("broadcast/user/publish", broadcast.SendUsers)
		MessageRouter.POST("broadcast/online/publish", broadcast.SendOnlineUsers)
	}
}
