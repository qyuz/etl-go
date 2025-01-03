package main

import (
	"fmt"
	"testing"
)

func TestGenerateSessionId(t *testing.T) {
	// Create request token at https://developer.themoviedb.org/reference/authentication-create-request-token
	// Grant request token access at https://www.themoviedb.org/authenticate/<request_token>/allow
	requestToken := "asdf"

	tmdbApi := getTmdbApi()

	if requestToken != "" {
		sessionResponse, err := tmdbApi.GetAuthSession(requestToken)
		if err != nil || !sessionResponse.Success {
			panic("Grant access to the application in the browser to get the request token, navigate to https://www.themoviedb.org/authenticate/<request_token>/allow")
		}
		fmt.Printf("Session ID: %s\n", sessionResponse.SessionID)
		panic("Session ID generated. Remove requestToken and re-run test")
	}
}

func TestGetWatchlistSeries(t *testing.T) {
	tmdbService := NewTmdbService(nil)
	tmdbSeries := tmdbService.GetWatchlistSeries()
	if len(tmdbSeries) == 0 {
		t.Error("Expected at least one series in watchlist")
	}
	for _, series := range tmdbSeries {
		fmt.Printf("ID: %d, Name: %s\n", series.ID, series.Name)
	}
}
