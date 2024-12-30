package load

import "testing"

var (
	notionPageServiceMock = &NotionPageServiceMock{}
)

func TestUpsertPageFoundMock(t *testing.T) {
	moviePage := &MoviePage{ID: "exists", notionPageService: notionPageServiceMock}
	moviePage.Upsert()
}

func TestUpsertPageNotFoundMock(t *testing.T) {
	moviePage := &MoviePage{ID: "does not exist", notionPageService: notionPageServiceMock}
	moviePage.Upsert()
}
