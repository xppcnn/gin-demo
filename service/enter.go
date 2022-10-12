package service

import (
	"github.com/xppcnn/gin-demo/service/article"
	"github.com/xppcnn/gin-demo/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ArticleServiceGroup article.ServiceGroup
}

var ServiceGroupVo = new(ServiceGroup)
