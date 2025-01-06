package load

type NotionService interface {
	UpsertVideoMedia(videoMedia VideoMedia)
}

type VideoMedia struct {
	Id   string
	Name string
}
