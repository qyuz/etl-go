package load

type NotionApiClientMock struct {
	CreateDatabasePageCalls       [][]Property
	CreateDatabasePageCalled      int
	QueryDatabasePageExistsCalls  []TextProperty
	QueryDatabasePageExistsCalled int
}

func (n *NotionApiClientMock) CreateDatabasePage(properties []Property) string {
	n.CreateDatabasePageCalled++

	n.CreateDatabasePageCalls = append(n.CreateDatabasePageCalls, properties)

	return ""
}

func (n *NotionApiClientMock) QueryDatabasePageExists(property TextProperty) bool {
	n.QueryDatabasePageExistsCalled++

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
