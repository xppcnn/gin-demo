package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xppcnn/gin-demo/config"
	"github.com/xppcnn/gin-demo/middleware"
	"github.com/xppcnn/gin-demo/pkg/setting"
	"github.com/xppcnn/gin-demo/routers/api/auth"
	"github.com/xppcnn/gin-demo/routers/business"
	"github.com/xppcnn/gin-demo/routers/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Business business.RouterGroup
}

var RouterGroupVo = new(RouterGroup)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type LoginForm struct {
	UserName string `json:"userName" binding:"required"`
	PassWord string `json:"passWord" binding:"required"`
}

func InitRouter() *gin.Engine {
	r := gin.New()
	conf := config.GetConfig()
	if err := middleware.InitLogger(conf.LogConfig, conf.RunMode); err != nil {
		fmt.Println(err)
	}
	// r.Use(middleware.GinLogger())
	// r.Use(middleware.GinRecovery(true))
	gin.SetMode(setting.RunMode)
	r.POST("/auth", auth.GetAuth)
	// apiv1 := r.Group("/api/v1").Use(middleware.JWT())
	// {
	// apiv1.GET("/tags", v1.GetTags)
	// apiv1.POST("/tags", v1.AddTag)
	// apiv1.PUT("/tags/:id", v1.EditTag)
	// apiv1.DELETE("/tags/:id", v1.DeleteTag)
	// apiv1.GET("/articles", v1.GetArticles)
	// apiv1.GET("/articles/:id", v1.GetArticle)
	// apiv1.POST("/articles", v1.AddArticle)
	// apiv1.PUT("/articles/:id", v1.EditArticle)
	// apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	// apiv1.POST("/register", v1.Register)
	// }
	return r
}
