package main

import (
	"github.com/ryanbradynd05/go-tmdb"
)

var (
	apiKey    = "2a5641277e811c5ab149ad870a6e7f2b"
	sessionId = "c710a8f348ed2d69e713ee59c204521a6775b9e5"
)

type TmdbService struct {
	tmdbApi *tmdb.TMDb
}

type TmdbSeries struct {
	ID   int
	Name string
}

func NewTmdbService(tmdbApi *tmdb.TMDb) TmdbService {
	if tmdbApi == nil {
		tmdbApi = getTmdbApi()
	}

	return TmdbService{
		tmdbApi: tmdbApi,
	}
}

func (t TmdbService) GetWatchlistSeries() []TmdbSeries {
	tvPageResults, err := t.tmdbApi.GetAccountWatchlistTv(21674720, sessionId, map[string]string{})
	if err != nil {
		panic(err)
	}

	tmdbSeries := []TmdbSeries{}
	for _, tv := range tvPageResults.Results {
		tmdbSeries = append(tmdbSeries, TmdbSeries{
			ID:   tv.ID,
			Name: tv.Name,
		})
	}

	return tmdbSeries
}

func getTmdbApi() *tmdb.TMDb {
	config := tmdb.Config{
		APIKey: apiKey,
	}
	return tmdb.Init(config)
}
