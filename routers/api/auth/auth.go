package auth

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/pkg/e"
	"github.com/xppcnn/gin-demo/utils"
	"go.uber.org/zap"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(ctx *gin.Context) {
	var authObj auth
	ctx.BindJSON(&authObj)
	valid := validation.Validation{}
	ok, _ := valid.Valid(&authObj)
	data := make(map[string]interface{})
	if ok {
		user, err := models.CheckAccount(authObj.Username, authObj.Password)
		if err == nil {
			idStr := user.ID.String()
			token, err := utils.GenerateToken(user.UserName, idStr)
			if err != nil {
				zap.L().Info(fmt.Sprintf("token生成失败:%s", err))
				e.FailWithMessage(err.Error(), ctx)
			} else {
				data["token"] = token
				e.OkWithDetailed(data, "toke生成成功", ctx)
			}
		} else {
			e.FailWithMessage(err.Error(), ctx)
		}
	} else {
		for _, err := range valid.Errors {
			zap.L().Info(fmt.Sprintf("%s:%s", err.Key, err.Message))
			e.BadWithMessage(err.Message, ctx)
			// logging.Info(err.Key, err.Message)
		}
	}
}
