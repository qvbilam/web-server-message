package resource

type UserObject struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Extra  string `json:"extra"`
}

// TextObject 文本消息
type TextObject struct {
	Content string     `json:"content"`
	User    UserObject `json:"user"`
	Extra   string     `json:"extra"`
}

// ImageObject 文件消息
type ImageObject struct {
	Content string     `json:"content"` // 缩略图
	Url     string     `json:"url"`
	User    UserObject `json:"user"`
	Extra   string     `json:"extra"`
}

// GIFObject GIF消息
type GIFObject struct {
	Url    string     `json:"url"`
	Width  int        `json:"width"`
	Height int        `json:"height"`
	Size   int        `json:"size"`
	User   UserObject `json:"user"`
	Extra  string     `json:"extra"`
}

// VoiceObject 音频消息
type VoiceObject struct {
	Url    string     `json:"url"`
	Second int        `json:"second"`
	User   UserObject `json:"user"`
	Extra  string     `json:"extra"`
}

// VideoObject 视频消息
type VideoObject struct {
	Name    string     `json:"name"`
	Content string     `json:"content"` // 缩略图
	Url     string     `json:"url"`
	Size    string     `json:"size"`
	Second  int        `json:"second"`
	User    UserObject `json:"user"`
	Extra   string     `json:"extra"`
}

// FileObject 文件消息
type FileObject struct {
	Name  string     `json:"name"`
	Type  string     `json:"type"`
	Size  int        `json:"size"`
	Url   string     `json:"url"`
	User  UserObject `json:"user"`
	Extra string     `json:"extra"`
}

// LBSObject 位置消息
type LBSObject struct {
	Content   string     `json:"content"` // 位置缩略图
	Latitude  string     `json:"latitude"`
	Longitude string     `json:"longitude"`
	Poi       string     `json:"poi"`
	User      UserObject `json:"user"`
	Extra     string     `json:"extra"`
}
