package business

import (
	"context"
	proto "message/api/qvbilam/message/v1"
	"message/global"
)

type PrivateBusiness struct {
	UserId       int64
	TargetUserId int64
	Message      *MessageBusiness
}

func (b *PrivateBusiness) Send() error {
	_, err := global.MessageServerClient.CreatePrivateMessage(context.Background(), &proto.CreatePrivateRequest{
		UserId:       b.UserId,
		TargetUserId: b.TargetUserId,
		Message: &proto.MessageRequest{
			Type:    b.Message.ContentType,
			Content: b.Message.Content,
			Url:     b.Message.Url,
			Extra:   b.Message.Extra,
		},
	})

	return err
}
