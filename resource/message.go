package resource

import proto "message/api/qvbilam/message/v1"

type MessagesResource struct {
	Total int64              `json:"total"`
	List  []*MessageResource `json:"list"`
}

type MessageResource struct {
	UserId int64  `json:"user_id"`
	UID    string `json:"uid"`
	Type   string `json:"type"`
	//Introduce string `json:"introduce"`
	Content     string `json:"content"`
	CreatedTime int64  `json:"created_time"`
}

func (r *MessagesResource) Resource(p *proto.MessagesResponse) *MessagesResource {
	m := MessagesResource{}
	m.Total = p.Total
	var ms []*MessageResource

	for _, item := range p.Messages {
		ms = append(ms, &MessageResource{
			UserId:      item.UserId,
			UID:         item.UID,
			Type:        item.Type,
			Content:     item.Content,
			CreatedTime: item.CreatedTime,
		})
	}

	m.List = ms
	return &m
}
