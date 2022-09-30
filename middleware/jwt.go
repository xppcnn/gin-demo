package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xppcnn/gin-demo/models"
	"github.com/xppcnn/gin-demo/pkg/e"
	"github.com/xppcnn/gin-demo/utils"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var token string
		var data interface{}
		code = e.SUCCESS
		bearerToken := ctx.GetHeader("Authorization")

		if strings.Contains(bearerToken, "Bearer ") {
			token = strings.Split(bearerToken, "Bearer ")[1]
		} else {
			token = bearerToken
		}
		if token == "" {
			code = e.ERROR_AUTH
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			//查询用户是否存在
			if _, err := models.FindUserById(claims.ID); err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
		}
		if code != e.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": code, "msg": e.GetMsg(code), "data": data})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
