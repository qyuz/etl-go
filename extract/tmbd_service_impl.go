package extract

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ryanbradynd05/go-tmdb"
)

type TmdbServiceImpl struct {
	bearer  string
	tmdbApi *tmdb.TMDb
}

func NewTmdbService(apiKey string, bearer string) TmdbService {
	config := tmdb.Config{
		APIKey: apiKey,
	}
	tmdbApi := tmdb.Init(config)

	return TmdbServiceImpl{
		bearer:  bearer,
		tmdbApi: tmdbApi,
	}
}

type GetWatchlistSeriesResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool     `json:"adult"`
		BackdropPath     string   `json:"backdrop_path"`
		GenreIDS         []int    `json:"genre_ids"`
		ID               int      `json:"id"`
		OriginCountry    []string `json:"origin_country"`
		OriginalLanguage string   `json:"original_language"`
		OriginalName     string   `json:"original_name"`
		Overview         string   `json:"overview"`
		Popularity       float64  `json:populatiry"`
		PosterPath       string   `json:"poster_path"`
		//	FirstAirDate     time.Time `json:"first_air_date"`
		Name        string  `json:"name"`
		VoteAverage float64 `json:"vote_average"`
		VoteCount   int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

func (t TmdbServiceImpl) GetWatchlistSeries() []TmdbSeries {
	// Docs at https://developer.themoviedb.org/reference/account-watchlist-tv
	url := "https://api.themoviedb.org/3/account/21674720/watchlist/tv?language=en-US&page=1&sort_by=created_at.asc"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.bearer))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//var prettyJSON bytes.Buffer
	//json.Indent(&prettyJSON, body, "", "\t")
	//fmt.Println(string(prettyJSON.String()))

	var watchlistSeriesResponse GetWatchlistSeriesResponse
	if err := json.Unmarshal(body, &watchlistSeriesResponse); err != nil {
		panic(err)
	}

	tmdbSeries := []TmdbSeries{}
	for _, series := range watchlistSeriesResponse.Results {
		tmdbSeries = append(tmdbSeries, TmdbSeries{
			Id:   series.ID,
			Name: series.Name,
		})
	}

	return tmdbSeries
}
