package pipeline

import (
	"etl/config"
	"etl/extract"
	"etl/load"
	"etl/transform"
	"fmt"
)

func InitAndRunTmdbWatchlistSeriesToNotionDbPipeline() {
	tmdbService := extract.NewTmdbService(config.TmdbApiKey, config.TmdbBearer)
	notionService := &load.NotionServiceImpl{}

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService, notionService)
}

func RunTmdbWatchlistSeriesToNotionDbPipeline(
	tmdbService extract.TmdbService,
	notionService load.NotionService,
) {
	tmdbSeries := tmdbService.GetWatchlistSeries()
	for _, series := range tmdbSeries {
		videoMedia := transform.TransformTmdbSeriesToNotionVideoMedia(series)
		notionService.UpsertVideoMedia(videoMedia)

		fmt.Println("Upserting Notion page for series: ", videoMedia.Name)
	}
}
