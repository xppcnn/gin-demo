package v1

import (
	"github.com/xppcnn/gin-demo/api/v1/article"
	"github.com/xppcnn/gin-demo/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ArticleApiGroup article.ApiGroup
}

var ApiGroupVo = new(ApiGroup)
