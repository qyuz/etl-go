package load

import (
	"fmt"

	"github.com/jomei/notionapi"
)

type NotionPageServiceMock struct {
}

func (n *NotionPageServiceMock) CheckPageExists(id string) bool {
	if id == "exists" {
		fmt.Println("Page found")
		return true
	}
	fmt.Println("Page not found")
	return false
}

func (n *NotionPageServiceMock) CreateDatabasePage() *notionapi.Page {
	fmt.Println("Creating new page")
	return nil
}
