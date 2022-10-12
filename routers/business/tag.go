package business

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xppcnn/gin-demo/api/v1"
)

type TagRouter struct {
}

func (a *TagRouter) InitTagRouter(Router *gin.RouterGroup) {
	apiVo := v1.ApiGroupVo.ArticleApiGroup.TagApi
	{
		Router.POST("/tags", apiVo.AddTag)
		Router.GET("/tags", apiVo.GetTags)
		Router.PUT("/tags/:id", apiVo.EditTag)
		Router.DELETE("/tags/:id", apiVo.DeleteTag)
	}
}
