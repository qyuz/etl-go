package extract

type TmdbServiceMock struct {
	GetWatchlistSeriesMock *func() []TmdbSeries
}

func (t *TmdbServiceMock) GetWatchlistSeries() []TmdbSeries {
	if t.GetWatchlistSeriesMock != nil {
		return (*t.GetWatchlistSeriesMock)()
	}
	return GetWatchlistSeriesMockData()
}

func GetWatchlistSeriesMockData() []TmdbSeries {
	tmdbSeries := []TmdbSeries{
		{ID: 1, Name: "Breaking Bad"},
	}
	return tmdbSeries
}
