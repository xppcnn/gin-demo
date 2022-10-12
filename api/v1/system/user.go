package system

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/pkg/e"
	"go.uber.org/zap"
)

type UserApi struct{}

func (u *UserApi) Register(ctx *gin.Context) {
	var vo = new(models.User)
	ctx.BindJSON(vo)
	valid := validation.Validation{}
	valid.Required(vo.UserName, "user_name").Message("用户名不得为空")
	valid.Required(vo.PassWord, "pass_word").Message("密码不得为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			e.BadWithMessage(err.Message, ctx)
			return
		}
	}
	user, err := userService.Register(*vo)
	if err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		e.FailWithDetailed(user, err.Error(), ctx)
	} else {
		zap.L().Info("注册成功")
		e.OkWithDetailed(user, "注册成功", ctx)
	}
}
