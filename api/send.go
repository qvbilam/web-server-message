package api

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "message/api/qvbilam/message/v1"
	"message/global"
	"message/validate"
)

func GroupSend(ctx *gin.Context) {
	uID, _ := ctx.Get("userId")
	userID := uID.(int64)

	request := validate.GroupValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		HandleValidateError(ctx, err)
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
		HandleGrpcErrorToHttp(ctx, err)
		return
	}
	SuccessNotContent(ctx)
}
