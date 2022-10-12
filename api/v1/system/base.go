package system

import (
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/pkg/e"
	"github.com/xppcnn/gin-demo/utils"
	"go.uber.org/zap"
)

type BaseApi struct {
}

func (b *BaseApi) Login(ctx *gin.Context) {
	var vo models.LoginForm
	ctx.ShouldBindJSON(&vo)
	valid := validation.Validation{}
	valid.Required(vo.UserName, "user_name").Message("用户名不得为空")
	valid.Required(vo.PassWord, "pass_word").Message("密码不得为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			e.BadWithMessage(err.Message, ctx)
			return
		}
	}
	user, err := userService.CheckUserByPassWord(vo)
	if err != nil {
		e.FailWithMessage(err.Error(), ctx)
		return
	}
	b.GetAuth(ctx, *user)
	// e.OkWithDetailed(user, "登录成功", ctx)
}

func (b *BaseApi) GetAuth(ctx *gin.Context, user models.PublicUser) {
	data := make(map[string]interface{})
	idStr := user.ID.String()
	token, err := utils.GenerateToken(user.UserName, idStr)
	if err != nil {
		zap.L().Info(fmt.Sprintf("token生成失败:%s", err))
		e.FailWithMessage(err.Error(), ctx)
	} else {
		data["token"] = token
		data["user"] = user
		e.OkWithDetailed(data, "登录成功", ctx)
	}
}
