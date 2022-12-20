package validate

type PrivateValidate struct {
	TargetUserId int64  `form:"target_user_id" json:"target_user_id" binding:"required"`
	ContentType  string `form:"content_type" json:"content_type" binding:"required"`
	Content      string `form:"content" json:"content" binding:"required"`
	Url          string `form:"url" json:"url" binding:"omitempty,url"`
	Extra        string `form:"extra" json:"extra" binding:"omitempty"`
}
