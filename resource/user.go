package resource

type UserResource struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Avatar string `json:"avatar"`
}
