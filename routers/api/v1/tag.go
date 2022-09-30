package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/utils"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/xppcnn/gin-demo/pkg/e"
	"github.com/xppcnn/gin-demo/pkg/setting"
)

// 获取文章多个tag
func GetTags(ctx *gin.Context) {
	name := ctx.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := ctx.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(utils.GetPage(ctx), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
}

func AddTag(ctx *gin.Context) {
	maps := make(map[string]string)
	ctx.BindJSON(&maps)
	name := maps["name"]
	state := com.StrTo(maps["state"]).MustInt()
	createdBy := maps["created_by"]
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不得为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditTag(ctx *gin.Context) {
	maps := make(map[string]string)
	ctx.BindJSON(&maps)

	id := ctx.Param("id")
	name := maps["name"]
	state := com.StrTo(maps["state"]).MustInt()
	modifiedBy := maps["modified_by"]
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(name, "name").Message("名称不得为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": make(map[string]string), "msg": e.GetMsg(code)})
}

func DeleteTag(ctx *gin.Context) {
	id := ctx.Param("id")
	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
