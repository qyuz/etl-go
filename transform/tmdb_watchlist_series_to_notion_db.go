package transform

import (
	"etl/extract"
	"etl/load"
)

func TransformTmdbWatchlistSeriesToNotionDb(tmdbSeries extract.TmdbSeries) *load.Movie {
	return &load.Movie{
		ID:   string(tmdbSeries.ID),
		Name: tmdbSeries.Name,
	}
}
