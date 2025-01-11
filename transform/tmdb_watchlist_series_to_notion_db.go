package transform

import (
	"etl/extract"
	"etl/load"
	"strconv"
)

func TransformTmdbSeriesToNotionVideoMedia(tmdbSeries extract.TmdbSeries) load.VideoMedia {
	return load.VideoMedia{
		Id:   strconv.Itoa(tmdbSeries.ID),
		Name: tmdbSeries.Name,
	}
}
