package extract

type TmdbServiceMock struct {
	GetWatchlistSeriesData []TmdbSeries
}

type TmdbServiceMockOpts struct {
	GetWatchlistSeriesData *[]TmdbSeries
}

func NewTmdbServiceMock(tmdbServiceMockOpts ...TmdbServiceMockOpts) TmdbServiceMock {
	var mockOpts *TmdbServiceMockOpts
	var GetWatchlistSeriesData *[]TmdbSeries

	if len(tmdbServiceMockOpts) > 0 {
		mockOpts = &tmdbServiceMockOpts[0]
	}

	if mockOpts == nil || mockOpts.GetWatchlistSeriesData == nil {
		GetWatchlistSeriesData = TmdbServiceMock{}.GetWatchlistSeriesDefaultData()
	} else {
		GetWatchlistSeriesData = mockOpts.GetWatchlistSeriesData
	}

	return TmdbServiceMock{GetWatchlistSeriesData: *GetWatchlistSeriesData}
}

func (t *TmdbServiceMock) GetWatchlistSeries() []TmdbSeries {
	return t.GetWatchlistSeriesData
}

func (TmdbServiceMock) GetWatchlistSeriesDefaultData() *[]TmdbSeries {
	tmdbSeries := []TmdbSeries{
		{ID: 1, Name: "Game of Thrones"},
		{ID: 2, Name: "Breaking Bad"},
	}
	return &tmdbSeries
}
