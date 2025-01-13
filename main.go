package main

import "etl/pipeline"

func main() {
	pipeline.InitAndRunTmdbWatchlistSeriesToNotionDbPipeline()
}
