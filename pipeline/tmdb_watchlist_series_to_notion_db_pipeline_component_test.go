package pipeline

import (
	"etl/extract"
	"etl/load"
	"fmt"
	"strconv"
	"testing"
)

func TestRunTmdbWatchlistSeriesToNotionDbPipelineWhenPageDoesntExistShouldCreateNewPage(t *testing.T) {
	tmdbServiceMock := extract.TmdbServiceMock{}

	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionServiceImpl{NotionApiClient: notionApiClientMock}

	// Act
	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	// Assert
	if notionApiClientMock.QueryDatabasePageExistsCalled != 1 {
		panic(fmt.Sprintf("Expected QueryDatabasePageExists to be called 1 time, but was called %d times", notionApiClientMock.QueryDatabasePageExistsCalled))
	}

	if notionApiClientMock.CreateDatabasePageCalled != 1 {
		panic(fmt.Sprintf("Expected CreateDatabasePage to be called 1 time, but was called %d times", notionApiClientMock.CreateDatabasePageCalled))
	}
}

func TestRunTmdbWatchlistSeriesToNotionDbPipelineWhenPageExistShouldNotCreateNewPage(t *testing.T) {
	tmdbServiceMock := extract.TmdbServiceMock{}

	watchlistSeriesData := extract.GetWatchlistSeriesMockData()
	notionApiClientMock := &load.NotionApiClientMock{
		CreateDatabasePageCalls: [][]load.Property{
			{
				load.TextProperty{
					Name:  "Movie Id",
					Value: strconv.Itoa(watchlistSeriesData[0].ID),
				},
			},
		},
	}
	notionService := &load.NotionServiceImpl{NotionApiClient: notionApiClientMock}

	// Act
	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	// Assert
	if notionApiClientMock.QueryDatabasePageExistsCalled != 1 {
		panic(fmt.Sprintf("Expected QueryDatabasePageExists to be called 1 time, but was called %d times", notionApiClientMock.QueryDatabasePageExistsCalled))
	}

	if notionApiClientMock.CreateDatabasePageCalled != 0 {
		panic(fmt.Sprintf("Expected CreateDatabasePage to be called 0 time, but was called %d times", notionApiClientMock.CreateDatabasePageCalled))
	}
}

func TestRunTmdbWatchlistSeriesToNotionDbPipelineShouldUpsertAllProps(t *testing.T) {
	// TODO: Implement this test
	tmdbServiceMock := extract.TmdbServiceMock{}
	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionServiceImpl{NotionApiClient: notionApiClientMock}

	// Act
	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	// Assert
}
