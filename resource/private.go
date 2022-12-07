package resource

import "encoding/json"

type PrivateObject struct {
	SendUserId string      `json:"send_user_id"`
	TargetId   string      `json:"target_id"`
	ObjectName string      `json:"object_name"`
	Content    interface{} `json:"content"`
}

func (o *PrivateObject) Resource() []byte {
	body, _ := json.Marshal(o)
	return body
}
