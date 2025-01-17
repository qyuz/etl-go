package pipeline

import (
	"etl/extract"
	"etl/load"
	"fmt"
	"strconv"
	"testing"
)

func TestRunTmdbWatchlistSeriesToNotionDbPipeline_WhenPageDoesntExistShouldCreateNewPage(t *testing.T) {
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

func TestRunTmdbWatchlistSeriesToNotionDbPipeline_WhenPageExistShouldNotCreateNewPage(t *testing.T) {
	tmdbServiceMock := extract.TmdbServiceMock{}

	watchlistSeriesData := extract.GetWatchlistSeriesMockData()
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

func ExampleRunTmdbWatchlistSeriesToNotionDbPipeline_ShouldUpsertAllProps() {
	tmdbServiceMock := extract.TmdbServiceMock{}

	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionServiceImpl{NotionApiClient: notionApiClientMock}

	// Act
	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	// Assert
	createdDatabasePageProps := notionApiClientMock.CreateDatabasePageCalls[0]

	movieIdProperty, ok := createdDatabasePageProps[0].(load.TextProperty)
	fmt.Println("TextProperty:", ok, "| Name:", movieIdProperty.GetName(), "| Value:", movieIdProperty.Value)

	movieNameProperty, ok := createdDatabasePageProps[1].(load.TextProperty)
	fmt.Println("TextProperty:", ok, "| Name:", movieNameProperty.GetName(), "| Value:", movieNameProperty.Value)

	// Output:
	// TextProperty: true | Name: Movie ID | Value: 1
	// TextProperty: true | Name: Name | Value: Breaking Bad
}
