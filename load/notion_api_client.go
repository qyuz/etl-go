package load

type NotionApiClient interface {
	CreateDatabasePage(properties []Property) string
	QueryDatabasePageExists(property TextProperty) bool
}

type Property interface {
	GetName() string
	GetType() string
}

type TextProperty struct {
	Name  string
	Value string
}

func (p TextProperty) GetName() string {
	return p.Name
}

func (p TextProperty) GetType() string {
	return "text"
}
