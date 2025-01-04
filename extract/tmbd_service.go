package main

type TmdbService interface {
	GetWatchlistSeries() []TmdbSeries
}

type TmdbSeries struct {
	ID   int
	Name string
}
