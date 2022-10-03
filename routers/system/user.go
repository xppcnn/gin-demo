package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xppcnn/gin-demo/api/v1"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user")
	apiVo := v1.ApiGroupVo.SystemApiGroup.UserApi
	{
		userRouter.POST("/register", apiVo.Register)
	}

}
