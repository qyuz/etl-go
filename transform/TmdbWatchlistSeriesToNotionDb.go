package main

import (
	"etl/extract"
	"etl/load"
)

func transformTmdbWatchlistSeriesToNotionDb(tmdbSeries extract.TmdbSeries) *load.Movie {
	return &load.Movie{
		ID:   tmdbSeries.ID,
		Name: tmdbSeries.Name,
	}
}
