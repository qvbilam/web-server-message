package broadcast

import (
	"github.com/gin-gonic/gin"
	"message/api"
	"message/business"
	"message/validate"
)

func SendUsers(ctx *gin.Context) {
	request := validate.BroadcastUserValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	mb := business.MessageBusiness{
		Code:        request.Code,
		ContentType: request.ContentType,
		Content:     request.Content,
		Url:         request.Url,
		Extra:       request.Extra,
	}
	b := business.BroadcastUserBusiness{
		UserIds:         request.UserIds,
		MessageBusiness: &mb,
	}

	var err error
	if request.UserIds == nil {
		err = b.SendAll()
	} else {
		err = b.SendOnline()
	}

	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func SendOnlineUsers(ctx *gin.Context) {
	request := validate.BroadcastOnlineUserValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	mb := business.MessageBusiness{
		Code:        request.Code,
		ContentType: request.ContentType,
		Content:     request.Content,
		Url:         request.Url,
		Extra:       request.Extra,
	}
	b := business.BroadcastUserBusiness{
		MessageBusiness: &mb,
	}

	if err := b.SendOnline(); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
