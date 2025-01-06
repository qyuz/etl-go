package transform

import (
	"etl/extract"
	"etl/load"
)

func TransformTmdbSeriesToNotionVideoMedia(tmdbSeries extract.TmdbSeries) load.VideoMedia {
	return load.VideoMedia{
		Id:   string(tmdbSeries.ID),
		Name: tmdbSeries.Name,
	}
}
