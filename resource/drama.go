package resource

import videoV1 "video/api/qvbilam/video/v1"

type DramaResource struct {
	Id           int64              `json:"id"`
	Name         string             `json:"name"`
	Introduce    string             `json:"introduce"`
	Cover        string             `json:"cover"`
	User         *UserResource      `json:"user"`
	Episode      *[]EpisodeResource `json:"episode"`
	Score        float64            `json:"score"`
	NewEpisode   int64              `json:"new_episode"`
	PlayCount    int64              `json:"play_count"`
	BarrageCount int64              `json:"barrage_count"`
	TotalCount   int64              `json:"total_count"`
	IsNew        bool               `json:"is_new"`
	IsHot        bool               `json:"is_hot"`
	IsEnd        bool               `json:"is_end"`
	CreatedTime  int64              `json:"created_time"`
}

func (r *DramaResource) Resource(response *videoV1.DramaResponse) *DramaResource {
	d := DramaResource{
		Id:           response.Id,
		Name:         response.Name,
		Introduce:    response.Introduce,
		Cover:        response.Cover,
		User:         nil,
		Score:        float64(response.Score),
		NewEpisode:   response.NewEpisode,
		PlayCount:    response.PlayCount,
		BarrageCount: response.BarrageCount,
		TotalCount:   response.TotalCount,
		IsNew:        response.IsNew,
		IsHot:        response.IsHot,
		IsEnd:        response.IsEnd,
		CreatedTime:  response.CreatedTime,
	}
	var episodes []EpisodeResource
	for _, e := range response.Episode {
		episodes = append(episodes, EpisodeResource{
			Id:      e.Id,
			Episode: e.Episode,
			Video: VideoResource{
				Id:          e.Video.Id,
				FileId:      e.Video.FileId,
				Name:        e.Video.Name,
				Introduce:   e.Video.Introduce,
				Cover:       e.Video.Cover,
				CreatedTime: e.Video.CreatedTime,
			},
		})
	}
	d.Episode = &episodes
	return &d
}

func (r *DramaResource) Collection(response *videoV1.DramaListResponse) *[]DramaResource {
	var dramas []DramaResource
	for _, drama := range response.Drama {
		dramas = append(dramas, DramaResource{
			Id:        drama.Id,
			Name:      drama.Name,
			Introduce: drama.Introduce,
			Cover:     drama.Cover,
			User:      nil,
			//Episode:     &episodes,
			Score:        float64(drama.Score),
			NewEpisode:   drama.NewEpisode,
			PlayCount:    drama.PlayCount,
			BarrageCount: drama.BarrageCount,
			TotalCount:   drama.TotalCount,
			IsEnd:        drama.IsEnd,
			IsNew:        drama.IsNew,
			IsHot:        drama.IsHot,
			CreatedTime:  drama.CreatedTime,
		})
	}

	return &dramas
}
