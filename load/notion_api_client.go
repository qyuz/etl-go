package load

type NotionApiClienter interface {
	CreateDatabasePage(properties []Property) string
	QueryDatabasePageExists(property TextProperty) bool
}

type Property interface {
	GetName() string
}

type TextProperty struct {
	Name  string
	Value string
}

func (p TextProperty) GetName() string {
	return p.Name
}
