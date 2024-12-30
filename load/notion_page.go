package load

import "github.com/jomei/notionapi"

type NotionPageService interface {
	GetPageById(id string) *notionapi.Page
	CreateDatabasePage() *notionapi.Page
}
