package extract

type TmdbServiceMock struct {
	GetWatchlistSeriesData *[]TmdbSeries
	GetWatchlistSeriesMock *func() []TmdbSeries
}

type TmdbServiceMockOpts struct {
	GetWatchlistSeriesData *[]TmdbSeries
	GetWatchlistSeriesMock *func() []TmdbSeries
}

func NewTmdbServiceMock(tmdbServiceMockOpts ...TmdbServiceMockOpts) TmdbServiceMock {
	var mockOpts *TmdbServiceMockOpts
	var GetWatchlistSeriesData *[]TmdbSeries
	var GetWatchlistSeriesMock *func() []TmdbSeries

	if len(tmdbServiceMockOpts) > 0 {
		mockOpts = &tmdbServiceMockOpts[0]
	}

	if mockOpts == nil || mockOpts.GetWatchlistSeriesData == nil {
		GetWatchlistSeriesData = TmdbServiceMock{}.GetWatchlistSeriesDefaultData()
	} else {
		GetWatchlistSeriesData = mockOpts.GetWatchlistSeriesData
	}

	if mockOpts != nil {
		if mockOpts.GetWatchlistSeriesMock != nil {
			GetWatchlistSeriesMock = mockOpts.GetWatchlistSeriesMock
		}
	}

	return TmdbServiceMock{
		GetWatchlistSeriesData: GetWatchlistSeriesData,
		GetWatchlistSeriesMock: GetWatchlistSeriesMock,
	}
}

func (t *TmdbServiceMock) GetWatchlistSeries() []TmdbSeries {
	if t.GetWatchlistSeriesMock != nil {
		return (*t.GetWatchlistSeriesMock)()
	} else {
		return *TmdbServiceMock{}.GetWatchlistSeriesDefaultData()
	}
	return *t.GetWatchlistSeriesData
}

func (TmdbServiceMock) GetWatchlistSeriesDefaultData() *[]TmdbSeries {
	tmdbSeries := []TmdbSeries{
		{ID: 1, Name: "Game of Thrones"},
		{ID: 2, Name: "Breaking Bad"},
	}
	return &tmdbSeries
}
