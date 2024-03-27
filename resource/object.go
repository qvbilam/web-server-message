package resource

import (
	"encoding/json"
	"message/enum"
)

type PrivateObject struct {
	Type       string      `json:"type"`
	SendUserId int64       `json:"send_user_id"`
	TargetId   int64       `json:"target_id"`
	ObjectName string      `json:"object_name"`
	Content    interface{} `json:"content"`
}

func (o *PrivateObject) Encode() []byte {
	o.Type = enum.ObjTypeUser
	body, _ := json.Marshal(o)
	return body
}

func (o *PrivateObject) Decode(content []byte) PrivateObject {
	obj := PrivateObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return PrivateObject{}
	}
	return obj
}

type RoomObject struct {
	Type       string   `json:"type"`
	SendUserId string   `json:"send_user_id"`
	TargetId   string   `json:"target_id"`
	ObjectName string   `json:"object_name"`
	Content    struct{} `json:"content"`
}

func (o *RoomObject) Encode() []byte {
	body, _ := json.Marshal(o)
	return body
}

func (o *RoomObject) Decode(content []byte) RoomObject {
	obj := RoomObject{}
	err := json.Unmarshal(content, &obj)
	if err != nil {
		return RoomObject{}
	}
	return obj
}
