package load

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
		return
	}

	n.NotionApiClient.CreateDatabasePage(properties)
}
