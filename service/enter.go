package service

import "github.com/xppcnn/gin-demo/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupVo = new(ServiceGroup)
