package private

import (
	"context"
	"github.com/gin-gonic/gin"
	"message/api"
	proto "message/api/qvbilam/message/v1"
	pageProto "message/api/qvbilam/page/v1"
	"message/business"
	"message/enum"
	"message/global"
	"message/resource"
	"message/validate"
	"strconv"
)

func Message(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	// 接受用户
	targetId := ctx.Param("id")
	targetUserId, _ := strconv.ParseInt(targetId, 10, 64)

	request := validate.GetPrivateMessageValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	if request.Page <= 0 {
		request.Page = 1
	}
	if request.PerPage <= 0 {
		request.PerPage = 10
	}

	msg, err := global.MessageServerClient.GetPrivateMessage(context.Background(), &proto.GetPrivateMessageRequest{
		UserId:       userID,
		TargetUserId: targetUserId,
		Type:         request.Type,
		Keyword:      request.Keyword,
		Page: &pageProto.PageRequest{
			Page:    request.Page,
			PerPage: request.PerPage,
		},
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	res := resource.MessagesResource{}
	api.SuccessNotMessage(ctx, res.Resource(msg))
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

func Read(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.ReadMessageValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	_, err := global.MessageServerClient.ReadPrivateMessage(context.Background(), &proto.ReadPrivateMessageRequest{
		UserId:     userID,
		MessageUid: request.MessageUid,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}

	api.SuccessNotContent(ctx)
}

func Rollback(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.ReadMessageValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}

	_, err := global.MessageServerClient.RollbackMessage(context.Background(), &proto.RollbackMessageRequest{
		UserId:     userID,
		MessageUid: request.MessageUid,
		ObjectType: enum.ObjTypeUser,
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotContent(ctx)
}
