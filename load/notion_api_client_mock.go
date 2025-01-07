package load

type NotionApiClientMock struct {
	CreateDatabasePageCalls      [][]Property
	QueryDatabasePageExistsCalls []TextProperty
}

func (n *NotionApiClientMock) CreateDatabasePage(properties []Property) string {
	n.CreateDatabasePageCalls = append(n.CreateDatabasePageCalls, properties)
	return ""
}

func (n *NotionApiClientMock) QueryDatabasePageExists(property TextProperty) bool {
	n.QueryDatabasePageExistsCalls = append(n.QueryDatabasePageExistsCalls, property)

	for _, createDatabaseCall := range n.CreateDatabasePageCalls {
		for _, prop := range createDatabaseCall {
			if prop.GetName() == property.GetName() &&
				prop.(TextProperty).Value == property.Value {
				return true
			}
		}
	}

	return false
}
