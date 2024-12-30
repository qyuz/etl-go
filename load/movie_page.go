package load

import (
	"fmt"
)

type MoviePage struct {
	ID                string
	notionPageService NotionPageService
}

func (m *MoviePage) Upsert() {
	pageExists := m.notionPageService.CheckPageExists(m.ID)
	if pageExists {
		fmt.Println("Page not found, creating new page")
		m.notionPageService.CreateDatabasePage()
	} else {
		fmt.Println("Page found")
	}
}
