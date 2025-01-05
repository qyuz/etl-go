package pipeline

import (
	"etl/config"
	"etl/extract"
	"etl/transform"
	"fmt"
)

func InitAndRunTmdbWatchlistSeriesToNotionDbPipeline() {
	tmdbService := extract.NewTmdbService(config.TmdbApiKey, config.TmdbBearer)

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService)
}

func RunTmdbWatchlistSeriesToNotionDbPipeline(
	tmdbService extract.TmdbService,
) {
	tmdbSeries := tmdbService.GetWatchlistSeries()
	for _, series := range tmdbSeries {
		notionPage := transform.TransformTmdbWatchlistSeriesToNotionDb(series)
		fmt.Println("Creating Notion page for series: ", notionPage.Name)
	}
}
