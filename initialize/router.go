package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/xppcnn/gin-demo/middleware"
	"github.com/xppcnn/gin-demo/routers"
)

func InitRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	systemRouter := routers.RouterGroupVo.System
	PublicGroup := Router.Group("")
	{
		systemRouter.BaseRouter.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWT())

	{
		systemRouter.UserRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}
