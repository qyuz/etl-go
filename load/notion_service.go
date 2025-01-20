package load

type NotionServicer interface {
	UpsertVideoMedia(videoMedia VideoMedia)
}

type VideoMedia struct {
	Id   string
	Name string
}
