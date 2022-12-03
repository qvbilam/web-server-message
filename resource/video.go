package resource

type VideoResource struct {
	Id          int64  `json:"id"`
	FileId      int64  `json:"file_id"`
	Name        string `json:"name"`
	Introduce   string `json:"introduce"`
	Cover       string `json:"cover"`
	CreatedTime int64  `json:"created_time"`
}
