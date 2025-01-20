package pipeline

import (
	"etl/config"
	"etl/extract"
	"etl/load"
	"etl/transform"
)

func InitAndRunTmdbWatchlistSeriesToNotionDbPipeline(config *config.Config) {
	tmdbService := extract.NewTmdbService(config.TmdbApiKey, config.TmdbBearer)

	notionApiClient := load.NewNotionApiClient(config.NotionApiKey, config.NotionDatabaseId)
	notionService := load.NotionService{NotionApiClient: notionApiClient}

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService, &notionService)
}

func RunTmdbWatchlistSeriesToNotionDbPipeline(
	tmdbService extract.TmdbServicer,
	notionService load.NotionServicer,
) {
	tmdbSeries := tmdbService.GetWatchlistSeries()

	for _, series := range tmdbSeries {
		videoMedia := transform.TransformTmdbSeriesToNotionVideoMedia(series)

		notionService.UpsertVideoMedia(videoMedia)
	}
}
