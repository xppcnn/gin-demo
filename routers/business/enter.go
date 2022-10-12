package business

type RouterGroup struct {
	ArticleRouter
	TagRouter
}

var RouterGroupVo = new(RouterGroup)
