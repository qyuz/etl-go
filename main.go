package main

import (
	"etl/config"
	"etl/pipeline"
)

func main() {
	config := config.NewConfig(".env.prod")
	pipeline.InitAndRunTmdbWatchlistSeriesToNotionDbPipeline(config)
}
