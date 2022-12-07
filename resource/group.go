package resource

import "encoding/json"

type GroupObject struct {
	SendUserId string   `json:"send_user_id"`
	TargetId   string   `json:"target_id"`
	ObjectName string   `json:"object_name"`
	Content    struct{} `json:"content"`
}

func (o *GroupObject) Resource() []byte {
	body, _ := json.Marshal(o)
	return body
}
