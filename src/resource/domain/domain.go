package domain

import "time"

type ResourceListResponse struct {
	Success bool       `json:"success"`
	Data    []Resource `json:"data"`
}

type Resource struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	PublishedAt time.Time  `json:"published_at"`
	ChannelId   string     `json:"channelId"`
	ID          ID         `json:"id"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type ID struct {
	Kind    string
	VideoID string
}

type ResourceQuery struct {
	Query          string
	ResourceType   string // resource,channel,playlist
	OrderBy        string // date preferred
	PublishedAfter string
}
