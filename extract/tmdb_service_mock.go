package extract

type TmdbServiceMock struct {
	GetWatchlistSeriesMock *func() []TmdbSeries
}

func (t *TmdbServiceMock) GetWatchlistSeries() []TmdbSeries {
	if t.GetWatchlistSeriesMock != nil {
		return (*t.GetWatchlistSeriesMock)()
	}
	return GetWatchlistSeriesData()
}

func GetWatchlistSeriesData() []TmdbSeries {
	tmdbSeries := []TmdbSeries{
		{ID: 1, Name: "Game of Thrones"},
		{ID: 2, Name: "Breaking Bad"},
	}
	return tmdbSeries
}
