package extract

import "testing"

type TmdbService interface {
	GetWatchlistSeries() []TmdbSeries
}

type TmdbSeries struct {
	Id   int
	Name string
}

func verifyShouldRunIntegrationTests(t *testing.T) {
	if false {
		t.Skip("skipping integration test")
	}
}
