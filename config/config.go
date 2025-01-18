package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NotionApiKey     string
	NotionDatabaseId string
	TmdbApiKey       string
	TmdbBearer       string
}

func NewConfig(envPath string) *Config {
	if err := godotenv.Load(envPath); err != nil {
		panic(err)
	}

	return &Config{
		NotionApiKey:     os.Getenv("NOTION_API_KEY"),
		NotionDatabaseId: os.Getenv("NOTION_DATABASE_ID"),
		TmdbApiKey:       os.Getenv("TMDB_API_KEY"),
		TmdbBearer:       os.Getenv("TMDB_BEARER"),
	}
}
