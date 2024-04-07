package youtube

type SearchResourceList struct {
	PublishedAfter string
	Query          string
	ResourceType   string // resource,channel,playlist
	OrderBy        string // date preferred
}
