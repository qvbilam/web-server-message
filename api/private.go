package api

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "message/api/qvbilam/message/v1"
	"message/global"
	"message/validate"
)

func PrivateSend(ctx *gin.Context) {
	// todo 获取登陆用户id
	userId := 2

	request := validate.PrivateValidate{}
	if err := ctx.ShouldBind(&request); err != nil {
		HandleValidateError(ctx, err)
		return
	}
	_, err := global.MessageServerClient.CreatePrivateMessage(context.Background(), &proto.CreatePrivateRequest{
		UserId:       int64(userId),
		TargetUserId: request.TargetUserId,
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
