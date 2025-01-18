//go:build ignore
// +build ignore

package main

import (
	"etl/config"
	"etl/pipeline"
)

func main() {
	config := config.NewConfig(".env.test")
	pipeline.InitAndRunTmdbWatchlistSeriesToNotionDbPipeline(config)
}
