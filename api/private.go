package api

import (
	"github.com/gin-gonic/gin"
	"message/global"
	"message/validate"
)

func Send(ctx *gin.Context) {
	// todo 获取登陆用户id
	userId := 2
	request := validate.PrivateValidate{}
	if err := ctx.Bind(&request); err != nil {
		//api.HandleValidateError(ctx, err)
		return
	}
	global.MessageServerClient.CreatePrivateMessage()
}
