package extract

type TmdbServiceMock struct {
}

func (t *TmdbServiceMock) GetWatchlistSeries() []TmdbSeries {
	return []TmdbSeries{
		{ID: 1, Name: "Game of Thrones"},
		{ID: 2, Name: "Breaking Bad"},
	}
}
