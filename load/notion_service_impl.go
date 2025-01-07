package load

type NotionServiceImpl struct {
	NotionApiClient NotionApiClient
}

func (n *NotionServiceImpl) UpsertVideoMedia(videoMedia VideoMedia) {
	movieIdProperty := TextProperty{"Movie Id", videoMedia.Id}

	properties := []Property{
		movieIdProperty,
		TextProperty{"Name", videoMedia.Name},
	}

	pageExists := n.NotionApiClient.QueryDatabasePageExists(movieIdProperty)
	if pageExists {
		return
	}

	n.NotionApiClient.CreateDatabasePage(properties)
}
