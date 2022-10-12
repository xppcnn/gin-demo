package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xppcnn/gin-demo/api/v1"
)

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	apiVo := v1.ApiGroupVo.SystemApiGroup.BaseApi
	{
		Router.POST("login", apiVo.Login)
	}
}
