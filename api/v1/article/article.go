package article

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/pkg/e"
	"github.com/xppcnn/gin-demo/pkg/setting"
	"github.com/xppcnn/gin-demo/utils"
)

type ArticleApi struct{}

func (articleApi *ArticleApi) GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
func (articleApi *ArticleApi) GetArticles(ctx *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(maps["state"].(int), 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId string = ""
	if tagId = ctx.Query("tag_id"); tagId != "" {
		maps["tag_id"] = tagId
		valid.Required(tagId, "tag_id").Message("ID不能为空")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(utils.GetPage(ctx), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "测试重启",
		"data": data,
	})

}

func (articleApi *ArticleApi) AddArticle(ctx *gin.Context) {
	maps := make(map[string]interface{})
	ctx.BindJSON(&maps)
	log.Printf("maps:%v,", maps)
	var state = -1
	if maps["state"] != "" {
		state = int(maps["state"].(float64))
	}
	valid := validation.Validation{}
	valid.Required(maps["tag_id"], "tag_id").Message("标签ID不得为空")
	valid.Required(maps["title"], "title").Message("标题不能为空")
	valid.Required(maps["desc"], "desc").Message("简述不能为空")
	valid.Required(maps["content"], "content").Message("内容不能为空")
	valid.Required(maps["created_by"], "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistTagByID(maps["tag_id"].(string)) {
			maps["state"] = state
			models.AddArticle(maps)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]interface{}),
	})
}
func (articleApi *ArticleApi) EditArticle(ctx *gin.Context) {
	maps := make(map[string]interface{})
	ctx.BindJSON(&maps)
	id := ctx.Param("id")
	title := maps["title"]
	desc := maps["desc"]
	content := maps["content"]
	modifiedBy := maps["modified_by"]
	var state = -1
	if maps["state"] != "" {
		state = int(maps["state"].(float64))
	}
	valid := validation.Validation{}
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(id, "id").Message("id不得为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			maps["state"] = state
			models.EditArticle(id, maps)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
func (articleApi *ArticleApi) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不得为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
