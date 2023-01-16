package group

import (
	"context"
	"github.com/gin-gonic/gin"
	"message/api"
	proto "message/api/qvbilam/message/v1"
	"message/enum"
	"message/global"
	"message/validate"
)

func Send(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.GroupCustomValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	_, err := global.MessageServerClient.CreateGroupMessage(context.Background(), &proto.CreateGroupRequest{
		UserId:  userID,
		GroupId: request.GroupId,
		Message: &proto.MessageRequest{
			Type:    request.ContentType,
			Content: request.Content,
			Url:     request.Url,
			Extra:   request.Extra,
		},
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotContent(ctx)
}

func SendText(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.GroupValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	_, err := global.MessageServerClient.CreateGroupMessage(context.Background(), &proto.CreateGroupRequest{
		UserId:  userID,
		GroupId: request.GroupId,
		Message: &proto.MessageRequest{
			Type:    enum.MsgTypeTxt,
			Content: request.Content,
			Url:     request.Url,
			Extra:   request.Extra,
		},
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotContent(ctx)
}

func SendImage(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.GroupValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	_, err := global.MessageServerClient.CreateGroupMessage(context.Background(), &proto.CreateGroupRequest{
		UserId:  userID,
		GroupId: request.GroupId,
		Message: &proto.MessageRequest{
			Type:    enum.MsgTypeImg,
			Content: request.Content,
			Url:     request.Url,
			Extra:   request.Extra,
		},
	})
	if err != nil {
		api.HandleGrpcErrorToHttp(ctx, err)
		return
	}
	api.SuccessNotContent(ctx)
}
