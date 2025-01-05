package pipeline

import (
	"etl/extract"
	"etl/load"
	"etl/transform"
)

func RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService extract.TmdbService, notionPageService load.NotionPageService) {
	tmdbSeries := e.tmdbSeriesService.GetWatchlistSeries()
	notionPages := transform.TransformTmdbWatchlistSeriesToNotionDb(tmdbSeries)
}
