package private

import (
	"context"
	"github.com/gin-gonic/gin"
	"message/api"
	proto "message/api/qvbilam/message/v1"
	"message/business"
	"message/enum"
	"message/global"
	"message/validate"
)

func message(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	global.MessageServerClient.GetPrivateMessage(context.Background(), &proto.GetPrivateMessageRequest{
		UserId:       userID,
		TargetUserId: 0,
	})
}

func Send(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.PrivateCustomValidate{}
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
	b := business.PrivateBusiness{
		UserId:       userID,
		TargetUserId: request.UserId,
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
		TargetUserId: request.UserId,
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
		TargetUserId: request.UserId,
		Message:      &mb,
	}

	if err := b.Send(); err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}
