package pipeline

import (
	"etl/extract"
	"etl/load"
	"testing"
)

func TestRunTmdbWatchlistSeriesToNotionDbPipeline(t *testing.T) {
	tmdbServiceMock := &extract.TmdbServiceMock{}
	notionServiceMock := &load.NotionServiceMock{}

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbServiceMock, notionServiceMock)

	if len(notionServiceMock.UpsertVideoMediaCalls) != 2 {
		t.Error("Expected UpsertVideoMedia to be called twice")
	}
}
