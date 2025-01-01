package load

import "testing"

func TestUpsertPageFound(t *testing.T) {
	runIntegrationTests(t)
	moviePage := &MoviePage{ID: "16626ea29ff180f39493cb1425c8ba54", notionPageService: &NotionPageServiceImpl{}}
	moviePage.Upsert()
}

func TestUpsertPageNotFound(t *testing.T) {
	runIntegrationTests(t)
	moviePage := &MoviePage{ID: "15d26ea29ff180d1b838ed84816e867d", notionPageService: &NotionPageServiceImpl{}}
	moviePage.Upsert()
}
