package group

import (
	"context"
	"github.com/gin-gonic/gin"
	"message/api"
	proto "message/api/qvbilam/message/v1"
	pageProto "message/api/qvbilam/page/v1"
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
	groupId, _ := strconv.ParseInt(targetId, 10, 64)

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

	msg, err := global.MessageServerClient.GetGroupMessage(context.Background(), &proto.GetGroupMessageRequest{
		UserId:  userID,
		GroupId: groupId,
		Keyword: request.Keyword,
		Type:    request.Type,
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

	request := validate.GroupCustomValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		api.HandleValidateError(ctx, err)
		return
	}
	_, err := global.MessageServerClient.CreateGroupMessage(context.Background(), &proto.CreateGroupRequest{
		UserId:  userID,
		GroupId: request.GroupId,
		Message: &proto.MessageRequest{
			Code:    request.Code,
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
