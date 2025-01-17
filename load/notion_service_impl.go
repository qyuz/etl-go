package load

import "log"

type NotionServiceImpl struct {
	NotionApiClient NotionApiClient
}

func (n *NotionServiceImpl) UpsertVideoMedia(videoMedia VideoMedia) {
	movieIdProperty := TextProperty{
		Name:  "Movie ID",
		Value: videoMedia.Id,
	}

	properties := []Property{
		movieIdProperty,
		TextProperty{
			Name:  "Name",
			Value: videoMedia.Name,
		},
	}

	pageExists := n.NotionApiClient.QueryDatabasePageExists(movieIdProperty)
	if pageExists {
		log.Println("Page exists for series:", videoMedia.Name)
		return
	}

	log.Println("Creating page for series:", videoMedia.Name)
	n.NotionApiClient.CreateDatabasePage(properties)
}
