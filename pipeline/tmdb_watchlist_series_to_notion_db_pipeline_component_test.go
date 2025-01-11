package pipeline

import (
	"etl/extract"
	"etl/load"
	"fmt"
	"testing"
)

func TestRunTmdbWatchlistSeriesToNotionDbPipeline(t *testing.T) {
	tmdbServiceMock := extract.TmdbServiceMock{}

	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionServiceImpl{NotionApiClient: notionApiClientMock}

	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	if len(notionApiClientMock.QueryDatabasePageExistsCalls) != 2 {
		panic(fmt.Sprintf("Expected QueryDatabasePageExists to be called 2 time, but was called %d times", len(notionApiClientMock.QueryDatabasePageExistsCalls)))
	}

	if len(notionApiClientMock.CreateDatabasePageCalls) != 2 {
		panic(fmt.Sprintf("Expected CreateDatabasePage to be called 2 time, but was called %d times", len(notionApiClientMock.CreateDatabasePageCalls)))
	}
}
