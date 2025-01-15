package extract

import (
	"etl/config"
	"fmt"
	"testing"
)

func TestGetWatchlistSeries(t *testing.T) {
	verifyShouldRunIntegrationTests(t)

	tmdbService := NewTmdbService(config.TmdbApiKey, config.TmdbBearer)
	tmdbSeries := tmdbService.GetWatchlistSeries()
	if len(tmdbSeries) == 0 {
		t.Error("Expected at least one series in watchlist")
	}
	for _, series := range tmdbSeries {
		fmt.Printf("ID: %d, Name: %s\n", series.Id, series.Name)
	}
}
