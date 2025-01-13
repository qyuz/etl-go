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
					Name:  "Movie ID",
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
	tmdbServiceMock := extract.TmdbServiceMock{}

	notionApiClientMock := &load.NotionApiClientMock{}
	notionService := &load.NotionServiceImpl{NotionApiClient: notionApiClientMock}

	// Act
	RunTmdbWatchlistSeriesToNotionDbPipeline(&tmdbServiceMock, notionService)

	// Assert
	createdDatabasePageProps := notionApiClientMock.CreateDatabasePageCalls[0]

	movieIdProperty, ok := createdDatabasePageProps[0].(load.TextProperty)
	if !ok {
		panic(fmt.Sprintf("Expected Movie ID property to be TextProperty, but got %s", movieIdProperty))
	}
	if movieIdProperty.GetName() != "Movie ID" {
		panic(fmt.Sprintf("Expected Movie ID property to be created, but got %s", movieIdProperty.GetName()))
	}
	if movieIdProperty.Value != "1" {
		panic(fmt.Sprintf("Expected Movie ID property value to be 1, but got %s", movieIdProperty.Value))
	}

	movieNameProperty, ok := createdDatabasePageProps[1].(load.TextProperty)
	if !ok {
		panic(fmt.Sprintf("Expected Name property to be TextProperty, but got %s", movieNameProperty))
	}
	if movieNameProperty.GetName() != "Name" {
		panic(fmt.Sprintf("Expected Name property to be created, but got %s", movieNameProperty.GetName()))
	}
	if movieNameProperty.Value != "Breaking Bad" {
		panic(fmt.Sprintf("Expected Name property value to be Breaking Bad, but got %s", movieNameProperty.Value))
	}

}
