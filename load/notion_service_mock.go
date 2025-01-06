package load

type NotionServiceMock struct {
	UpsertVideoMediaCalls []VideoMedia
}

func (n *NotionServiceMock) UpsertVideoMedia(videoMedia VideoMedia) {
	n.UpsertVideoMediaCalls = append(n.UpsertVideoMediaCalls, videoMedia)
}
