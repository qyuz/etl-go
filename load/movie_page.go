package load

import (
	"fmt"
)

type MoviePage struct {
	ID                string
	notionPageService NotionPageService
}

func (m *MoviePage) Upsert() {
	pageExists := m.notionPageService.CheckDatabasePageExists("Movie ID", m.ID)

	if pageExists {
		fmt.Println("Page found")
		return
	}

	fmt.Println("Page not found, creating new page")

	titleProperty := TextProperty{Name: "Title", Text: "Movie Page"}
	movieIdProperty := TextProperty{Name: "Movie ID", Text: m.ID}
	m.notionPageService.CreateDatabasePage([]Property{titleProperty, movieIdProperty})
}
