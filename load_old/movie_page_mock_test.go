package load_old

import (
	"testing"
)

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

	if notionPageServiceMock.CreateDatabasePageResult != "Title:text,Movie ID:text," {
		t.Error("UpsertPageNotFoundMock test failed")
	}
}
