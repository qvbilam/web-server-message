package validate

type DramaCreate struct {
	RegionId        int64  `form:"region_id" json:"region_id" binding:"required,numeric"`
	CategoryId      int64  `form:"category_id" json:"category_id" binding:"required,numeric"`
	Name            string `form:"name" json:"name" binding:"required,min=2,max=100"`
	Introduce       string `form:"introduce" json:"introduce" binding:"omitempty,min=0,max=100"`
	Cover           string `form:"cover" json:"cover" binding:"omitempty,url"`
	HorizontalCover string `form:"horizontal_cover" json:"horizontal_cover" binding:"omitempty,url"`
	TotalCount      int64  `form:"total_count" json:"total_count" binding:"omitempty,min=0,max=999999"`
}

type DramaUpdate struct {
	RegionId        *int64  `form:"region_id" json:"region_id" binding:"omitempty,numeric"`
	CategoryId      *int64  `form:"category_id" json:"category_id" binding:"omitempty,numeric"`
	Name            *string `form:"name" json:"name" binding:"omitempty,min=5,max=100"`
	Introduce       *string `form:"introduce" json:"introduce" binding:"omitempty,min=0,max=100"`
	Cover           *string `form:"cover" json:"cover" binding:"omitempty,url"`
	HorizontalCover *string `form:"horizontal_cover" json:"horizontal_cover" binding:"omitempty,url"`
	TotalCount      *int64  `form:"total_count" json:"total_count" binding:"omitempty,min=0,max=999999"`
}

type DramaSearch struct {
	UserId           int64  `form:"user_id" json:"user_id" binding:"omitempty,numeric"`
	RegionId         int64  `form:"region_id" json:"region_id" binding:"omitempty,numeric"`
	CategoryId       int64  `form:"category_id" json:"category_id" binding:"omitempty,numeric"`
	Keyword          string `form:"keyword" json:"keyword" binding:"omitempty,min=1,max=100"`
	Sort             string `form:"sort" json:"sort" binding:"omitempty"`
	IsRecommend      *bool  `form:"is_recommend" json:"is_recommend" binding:"omitempty,boolean"`
	IsNew            *bool  `form:"is_new" json:"is_new" binding:"omitempty,boolean"`
	IsHot            *bool  `form:"is_hot" json:"is_hot" binding:"omitempty,boolean"`
	IsEnd            *bool  `form:"is_end" json:"is_end" binding:"omitempty,boolean"`
	IsVisible        *bool  `form:"is_visible" json:"is_visible" binding:"omitempty,boolean"`
	TotalCountMin    int64  `form:"total_count_min" json:"total_count_min" binding:"omitempty,numeric"`
	TotalCountMax    int64  `form:"total_count_max" json:"total_count_max" binding:"omitempty,numeric"`
	FavoriteCountMin int64  `form:"favorite_count_min" json:"favorite_count_min" binding:"omitempty,numeric"`
	FavoriteCountMax int64  `form:"favorite_count_max" json:"favorite_count_max" binding:"omitempty,numeric"`
	LikeCountMin     int64  `form:"like_count_min" json:"like_count_min" binding:"omitempty,numeric"`
	LikeCountMax     int64  `form:"like_count_max" json:"like_count_max" binding:"omitempty,numeric"`
	PlayCountMin     int64  `form:"play_count_min" json:"play_count_min" binding:"omitempty,numeric"`
	PlayCountMax     int64  `form:"play_count_max" json:"play_count_max" binding:"omitempty,numeric"`
	BarrageCountMin  int64  `form:"barrage_count_min" json:"barrage_count_min" binding:"omitempty,numeric"`
	BarrageCountMax  int64  `form:"barrage_count_max" json:"barrage_count_max" binding:"omitempty,numeric"`
	Page             int64  `form:"page" json:"page" binding:"omitempty,numeric,min=1"`
	PerPage          int64  `form:"per_page" json:"per_page" binding:"omitempty,numeric,min=1"`
}
