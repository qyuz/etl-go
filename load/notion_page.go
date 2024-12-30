package load

import "github.com/jomei/notionapi"

type NotionPageService interface {
	CheckPageExists(id string) bool
	CreateDatabasePage() *notionapi.Page
}
