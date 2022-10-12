package article

import "github.com/xppcnn/gin-demo/service"

type ApiGroup struct {
	ArticleApi
	TagApi
}

var (
	tagService = service.ServiceGroupVo.ArticleServiceGroup.TagService
)
