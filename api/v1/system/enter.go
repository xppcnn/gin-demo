package system

import "github.com/xppcnn/gin-demo/service"

type ApiGroup struct {
	UserApi
	BaseApi
}

var (
	userService = service.ServiceGroupVo.SystemServiceGroup.UserService
)
