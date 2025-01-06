package load_old

import (
	"fmt"
)

type MoviePage struct {
	ID                string
	notionPageService NotionPageService
}

func (m *MoviePage) Upsert() {
	property := TextProperty{Name: "Movie ID", Text: m.ID}
	pageExists := m.notionPageService.CheckDatabasePageExists(property)

	if pageExists {
		fmt.Println("Page found")
		return
	}

	fmt.Println("Page not found, creating new page")

	titleProperty := TextProperty{Name: "Title", Text: "Movie Page"}
	movieIdProperty := TextProperty{Name: "Movie ID", Text: m.ID}
	m.notionPageService.CreateDatabasePage([]Property{titleProperty, movieIdProperty})
}
