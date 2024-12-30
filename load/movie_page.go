package load

import (
	"fmt"
)

type MoviePage struct {
	ID                string
	notionPageService NotionPageService
}

func (m *MoviePage) Upsert() {
	notionPage := &NotionPageServiceImpl{}
	pageExists := notionPage.CheckPageExists(m.ID)
	if pageExists {
		fmt.Println("Page not found, creating new page")
		notionPage.CreateDatabasePage()
	} else {
		fmt.Println("Page found")
	}
}
