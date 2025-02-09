package extract

type TmdbServiceMock struct {
	GetWatchlistSeriesMock *func() []TmdbSeries
}

func (t *TmdbServiceMock) GetWatchlistSeries() []TmdbSeries {
	if t.GetWatchlistSeriesMock != nil {
		return (*t.GetWatchlistSeriesMock)()
	}
	panic("GetWatchlistSeriesMock is not set")
}
