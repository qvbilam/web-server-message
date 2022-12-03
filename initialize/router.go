package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"video/middleware"
	videoRouter "video/router"
)

func InitRouters() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	apiRouter := router.Group("")

	// 初始化基础组建路由
	videoRouter.InitVideoRouter(apiRouter)
	videoRouter.InitEpisodesRouter(apiRouter)
	videoRouter.InitRegionRouter(apiRouter)
	videoRouter.InitCategoryRouter(apiRouter)
	videoRouter.InitBarrageRouter(apiRouter)
	videoRouter.InitDramaRouter(apiRouter)

	return router
}
