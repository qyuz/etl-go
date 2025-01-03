package load

import (
	"fmt"
)

type NotionPageServiceMock struct {
	CreateDatabasePageResult string
}

func (n *NotionPageServiceMock) CheckDatabasePageExists(property TextProperty) bool {
	if property.GetName() == "Movie ID" && property.Text == "exists" {
		fmt.Println("Page found")
		return true
	}
	fmt.Println("Page not found")
	return false
}

func (n *NotionPageServiceMock) CreateDatabasePage(properties []Property) string {
	fmt.Println("Creating new page")

	result := ""
	for _, property := range properties {
		result += property.GetName() + ":" + property.GetType() + ","
	}

	n.CreateDatabasePageResult = result

	return result
}
