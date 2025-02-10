package extract

import (
	"etl/config"
	"fmt"
	"testing"
)

func TestGetWatchlistSeries(t *testing.T) {
	verifyShouldRunIntegrationTests(t)

	config := config.NewConfig("../.env.test")

	tmdbService := NewTmdbService(config.TmdbApiKey, config.TmdbBearer)
	tmdbSeries := tmdbService.GetWatchlistSeries()

	if len(tmdbSeries) == 0 {
		t.Error("Expected at least one series in watchlist")
	}

	for _, series := range tmdbSeries {
		fmt.Printf("ID: %d, Name: %s\n", series.Id, series.Name)
	}
}

func TestGetSeriesDetails(t *testing.T) {
	verifyShouldRunIntegrationTests(t)

	config := config.NewConfig("../.env.test")
	tmdbService := NewTmdbService(config.TmdbApiKey, config.TmdbBearer)

	//	tmdbDetails := tmdbService.GetSeriesDetails(4607) // Lost
	tmdbDetails := tmdbService.GetSeriesDetails(95557) // Invincible

	fmt.Printf("ID: %d, Name: %s Last Air: %s Next Air: %s\n", tmdbDetails.ID, tmdbDetails.Name, tmdbDetails.LastAirDate, tmdbDetails.NextEpisodeToAir.AirDate)
}
