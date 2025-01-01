package load

import (
	"fmt"
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

func (n *NotionPageServiceMock) CreateDatabasePage(id string) string {
	fmt.Println("Creating new page")
	return "new_page_id"
}
