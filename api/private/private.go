package private

import (
	"github.com/gin-gonic/gin"
	"message/api"
	"message/business"
	"message/enum"
	"message/validate"
)

func Send(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.PrivateCustomValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	mb := business.MessageBusiness{
		ContentType: request.ContentType,
		Content:     request.Content,
		Url:         request.Url,
		Extra:       request.Extra,
	}
	b := business.PrivateBusiness{
		UserId:       userID,
		TargetUserId: request.TargetUserId,
		Message:      &mb,
	}

	if err := b.Send(); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func SendText(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.PrivateValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	mb := business.MessageBusiness{
		ContentType: enum.MsgTypeTxt,
		Content:     request.Content,
		Url:         request.Url,
		Extra:       request.Extra,
	}
	b := business.PrivateBusiness{
		UserId:       userID,
		TargetUserId: request.TargetUserId,
		Message:      &mb,
	}

	if err := b.Send(); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func SendImage(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.PrivateValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	mb := business.MessageBusiness{
		ContentType: enum.MsgTypeImg,
		Content:     request.Content, // 缩略图
		Url:         request.Url,     // 图片地址
		Extra:       request.Extra,
	}
	b := business.PrivateBusiness{
		UserId:       userID,
		TargetUserId: request.TargetUserId,
		Message:      &mb,
	}

	if err := b.Send(); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
