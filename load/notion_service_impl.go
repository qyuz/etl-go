package load

import "fmt"

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
		fmt.Println("Page exists for series: ", videoMedia.Name)
		return
	}

	fmt.Println("Creating page for series: ", videoMedia.Name)
	n.NotionApiClient.CreateDatabasePage(properties)
}
