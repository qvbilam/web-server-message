package initialize

import (
	"github.com/gin-gonic/gin"
	"message/api/ws"
	"message/middleware"
	messageRouter "message/router"
	"net/http"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})

	// websocket
	router.GET("/ws", ws.Handel)

	apiRouter := router.Group("")
	messageRouter.InitPrivateRouter(apiRouter)

	return router
}
