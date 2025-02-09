package pipeline

import (
	"etl/extract"
	"etl/load"
	"fmt"
	"strconv"
	"testing"
)

var getWatchlistSeriesMock = func() []extract.TmdbSeries {
	return []extract.TmdbSeries{
		{Id: 1, Name: "Breaking Bad"},
	}
}

func TestRunTmdbWatchlistSeriesToNotionDbPipeline_WhenPageDoesntExistShouldCreateNewPage(t *testing.T) {
	tmdbServiceMock := extract.TmdbServiceMock{
		GetWatchlistSeriesMock: &getWatchlistSeriesMock,
	}

	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionService{NotionApiClient: notionApiClientMock}

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

func TestRunTmdbWatchlistSeriesToNotionDbPipeline_WhenPageExistShouldNotCreateNewPage(t *testing.T) {
	tmdbServiceMock := extract.TmdbServiceMock{
		GetWatchlistSeriesMock: &getWatchlistSeriesMock,
	}

	watchlistSeriesData := getWatchlistSeriesMock()
	notionApiClientMock := &load.NotionApiClientMock{
		CreateDatabasePageCalls: [][]load.Property{
			{
				load.TextProperty{
					Name:  "Movie ID",
					Value: strconv.Itoa(watchlistSeriesData[0].Id),
				},
			},
		},
	}
	notionService := &load.NotionService{NotionApiClient: notionApiClientMock}

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

func ExampleRunTmdbWatchlistSeriesToNotionDbPipeline_ShouldUpsertAllProps() {
	tmdbServiceMock := extract.TmdbServiceMock{
		GetWatchlistSeriesMock: &getWatchlistSeriesMock,
	}

	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionService{NotionApiClient: notionApiClientMock}

	// Act
	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	// Assert
	fmt.Printf("%#v\n", notionApiClientMock.CreateDatabasePageCalls)

	// Output:
	// [][]load.Property{[]load.Property{load.TextProperty{Name:"Movie ID", Value:"1"}, load.TextProperty{Name:"Name", Value:"Breaking Bad"}}}
}
