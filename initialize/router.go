package initialize

import (
	"github.com/gin-gonic/gin"
	"message/api/ws"
	"message/middleware"
	messageRouter "message/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	// websocket
	router.GET("/ws", ws.Handel)

	apiRouter := router.Group("")
	messageRouter.InitMessageRouter(apiRouter)

	return router
}
