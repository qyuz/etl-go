package load

import (
	"fmt"

	"github.com/jomei/notionapi"
)

type MoviePage struct {
	ID         string
	notionPage *notionapi.Page
}

func Upsert(m *MoviePage, id string) {
	page := m.notionPage.GetPageById(id)
	if page == nil {
		fmt.Println("Page not found, creating new page")
		page = CreateDatabasePage()
	} else {
		fmt.Println("Page found")
	}
}
