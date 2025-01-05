package pipeline

import (
	"etl/extract"
	"testing"
)

func TestRunTmdbWatchlistSeriesToNotionDbPipeline(t *testing.T) {
	tmdbService := &extract.TmdbServiceMock{}

	RunTmdbWatchlistSeriesToNotionDbPipeline(tmdbService)
}
