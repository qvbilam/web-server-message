package validate

type BroadcastOnlineUserValidate struct {
	ContentType string `form:"content_type" json:"content_type" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Url         string `form:"url" json:"url" binding:"omitempty,url"`
	Extra       string `form:"extra" json:"extra" binding:"omitempty"`
	Code        int64  `form:"code" json:"code" binding:"omitempty"`
}

type BroadcastUserValidate struct {
	UserIds     []int64 `form:"user_ids" json:"user_ids" binding:"omitempty"`
	ContentType string  `form:"content_type" json:"content_type" binding:"required"`
	Content     string  `form:"content" json:"content" binding:"required"`
	Url         string  `form:"url" json:"url" binding:"omitempty,url"`
	Extra       string  `form:"extra" json:"extra" binding:"omitempty"`
	Code        int64   `form:"code" json:"code" binding:"omitempty"`
}

type SystemValidate struct {
	UserId  int64  `form:"user_id" json:"user_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Url     string `form:"url" json:"url" binding:"omitempty,url"`
	Extra   string `form:"extra" json:"extra" binding:"omitempty"`
}

type TipValidate struct {
	UserId  int64  `form:"user_id" json:"user_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Url     string `form:"url" json:"url" binding:"omitempty,url"`
	Extra   string `form:"extra" json:"extra" binding:"omitempty"`
}

type PrivateValidate struct {
	UserId  int64  `form:"user_id" json:"user_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Url     string `form:"url" json:"url" binding:"omitempty,url"`
	Extra   string `form:"extra" json:"extra" binding:"omitempty"`
}

type PrivateCustomValidate struct {
	UserId      int64  `form:"user_id" json:"user_id" binding:"required"`
	ContentType string `form:"content_type" json:"content_type" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Url         string `form:"url" json:"url" binding:"omitempty,url"`
	Code        int64  `form:"code" json:"code" binding:"omitempty"`
	Extra       string `form:"extra" json:"extra" binding:"omitempty"`
}

type GroupValidate struct {
	GroupId int64  `form:"group_id" json:"group_id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Url     string `form:"url" json:"url" binding:"omitempty,url"`
	Extra   string `form:"extra" json:"extra" binding:"omitempty"`
}

type GroupCustomValidate struct {
	GroupId     int64  `form:"group_id" json:"group_id" binding:"required"`
	ContentType string `form:"content_type" json:"content_type" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Url         string `form:"url" json:"url" binding:"omitempty,url"`
	Code        int64  `form:"code" json:"code" binding:"omitempty"`
	Extra       string `form:"extra" json:"extra" binding:"omitempty"`
}
