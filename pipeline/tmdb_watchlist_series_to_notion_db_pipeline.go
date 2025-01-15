package pipeline

import (
	"etl/config"
	"etl/extract"
	"etl/load"
	"etl/transform"
)

func InitAndRunTmdbWatchlistSeriesToNotionDbPipeline() {
	config := config.NewConfig()

	tmdbService := extract.NewTmdbService(config.TmdbApiKey, config.TmdbBearer)

	notionApiClient := load.NewNotionApiClientImpl(config.NotionApiKey, config.NotionDatabaseId)
	notionService := load.NotionServiceImpl{NotionApiClient: notionApiClient}

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService, &notionService)
}

func RunTmdbWatchlistSeriesToNotionDbPipeline(
	tmdbService extract.TmdbService,
	notionService load.NotionService,
) {
	tmdbSeries := tmdbService.GetWatchlistSeries()

	for _, series := range tmdbSeries {
		videoMedia := transform.TransformTmdbSeriesToNotionVideoMedia(series)

		notionService.UpsertVideoMedia(videoMedia)
	}
}
