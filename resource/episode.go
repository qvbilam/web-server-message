package resource

type EpisodeResource struct {
	Id      int64         `json:"id"`
	Episode int64         `json:"episode"`
	Video   VideoResource `json:"video"`
}
