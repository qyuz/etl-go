package extract

import "testing"

type TmdbService interface {
	GetWatchlistSeries() []TmdbSeries
}

type TmdbSeries struct {
	ID   int
	Name string
}

func verifyShouldRunIntegrationTests(t *testing.T) {
	if true {
		t.Skip("skipping integration test")
	}
}
