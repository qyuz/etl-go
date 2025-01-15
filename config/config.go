package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	NotionApiKey     string
	NotionDatabaseId string
	TmdbApiKey       string
	TmdbBearer       string
}

func NewConfig(path ...string) *config {
	var envPath string

	if len(path) > 0 {
		envPath = path[0]
	} else {
		envPath = ".env.test"
	}

	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	return &config{
		NotionApiKey:     os.Getenv("NOTION_API_KEY"),
		NotionDatabaseId: os.Getenv("NOTION_DATABASE_ID"),
		TmdbApiKey:       os.Getenv("TMDB_API_KEY"),
		TmdbBearer:       os.Getenv("TMDB_BEARER"),
	}
}
