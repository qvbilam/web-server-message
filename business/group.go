package business

import (
	"context"
	proto "message/api/qvbilam/message/v1"
	"message/global"
)

type GroupBusiness struct {
	UserId   int64
	TargetId int64
	Message  *MessageBusiness
}

func (b *GroupBusiness) Send() error {
	_, err := global.MessageServerClient.CreatePrivateMessage(context.Background(), &proto.CreatePrivateRequest{
		UserId:       b.UserId,
		TargetUserId: b.TargetId,
		Message: &proto.MessageRequest{
			Type:    b.Message.ContentType,
			Content: b.Message.Content,
			Url:     b.Message.Url,
			Extra:   b.Message.Extra,
		},
	})

	return err
}
