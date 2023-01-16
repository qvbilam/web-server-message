package business

import (
	"context"
	proto "message/api/qvbilam/message/v1"
	userProto "message/api/qvbilam/user/v1"
	"message/global"
)

type BroadcastUserBusiness struct {
	UserIds         []int64
	MessageBusiness *MessageBusiness
}

func (b *BroadcastUserBusiness) SendAll() error {
	users, err := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{})
	if err != nil {
		return err
	}
	if users == nil {
		return nil
	}
	var userIds []int64
	for _, u := range users.Users {
		userIds = append(userIds, u.Id)
	}

	b.UserIds = userIds
	return b.Send()
}

func (b *BroadcastUserBusiness) SendOnline() error {
	// todo 查找在线的
	users, err := global.UserServerClient.List(context.Background(), &userProto.SearchRequest{})
	if err != nil {
		return err
	}
	if users == nil {
		return nil
	}
	var userIds []int64
	for _, u := range users.Users {
		userIds = append(userIds, u.Id)
	}

	b.UserIds = userIds
	return b.Send()
}

func (b *BroadcastUserBusiness) Send() error {
	_, err := global.MessageServerClient.CreateUserBroadcastMessage(context.Background(), &proto.CreateUserBroadcastRequest{
		UserIds: b.UserIds,
		Message: &proto.MessageRequest{
			Code:    b.MessageBusiness.Code,
			Type:    b.MessageBusiness.ContentType,
			Content: b.MessageBusiness.Content,
			Url:     b.MessageBusiness.Url,
			Extra:   b.MessageBusiness.Extra,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
