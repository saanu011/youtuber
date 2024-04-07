package youtube

import "time"

type ResourceListResponse struct {
	Success bool       `json:"success"`
	Data    []Resource `json:"data"`
}

type Resource struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PublishedAt time.Time  `json:"published_at"`
	Thumbnails  Thumbnails `json:"thumbnails"`
	Statistics  Statistics `json:"statistics"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Statistics struct {
	ViewCount     string `json:"viewCount"`
	LikeCount     string `json:"likeCount"`
	DislikeCount  string `json:"dislikeCount"`
	FavoriteCount string `json:"favoriteCount"`
	CommentCount  string `json:"commentCount"`
}
