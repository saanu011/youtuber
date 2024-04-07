package domain

type Resource struct {
	PublishedAfter string
	Title          string
	Description    string
}

type ResourceQuery struct {
	Query          string
	ResourceType   string // resource,channel,playlist
	OrderBy        string // date preferred
	PublishedAfter string
}
