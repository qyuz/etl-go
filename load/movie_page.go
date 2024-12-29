package load

import (
	"fmt"

	"github.com/jomei/notionapi"
)

type MoviePage struct {
	ID         string
	notionPage *notionapi.Page
}

func (m *MoviePage) Upsert() {
	notionPage := &NotionPage{ID: m.ID}
	page := notionPage.GetPageById()
	if page == nil {
		fmt.Println("Page not found, creating new page")
		page = notionPage.CreateDatabasePage()
	} else {
		fmt.Println("Page found")
	}
}
