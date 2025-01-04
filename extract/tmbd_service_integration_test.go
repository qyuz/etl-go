package main

import (
	"fmt"
	"testing"
)

func TestGetWatchlistSeries(t *testing.T) {
	tmdbService := NewTmdbService()
	tmdbSeries := tmdbService.GetWatchlistSeries()
	if len(tmdbSeries) == 0 {
		t.Error("Expected at least one series in watchlist")
	}
	for _, series := range tmdbSeries {
		fmt.Printf("ID: %d, Name: %s\n", series.ID, series.Name)
	}
}
