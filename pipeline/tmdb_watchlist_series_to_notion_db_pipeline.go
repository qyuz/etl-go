package pipeline

import (
	"etl/config"
	"etl/extract"
	"etl/transform"
	"fmt"
)

func InitAndRunTmdbWatchlistSeriesToNotionDbPipeline() {
	tmdbService := extract.NewTmdbService(config.TmdbApiKey, config.TmdbBearer)

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService, notionPageService)
}

func RunTmdbWatchlistSeriesToNotionDbPipeline(
	tmdbService extract.TmdbService,
) {
	tmdbSeries := tmdbService.GetWatchlistSeries()
	notionPages := transform.TransformTmdbWatchlistSeriesToNotionDb(tmdbSeries)

	for _, notionPage := range notionPages {
		fmt.Println("Creating Notion page for series: ", notionPage.Name)
	}
}
