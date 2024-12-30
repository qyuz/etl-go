package load

import (
	"context"
	"strings"

	"github.com/jomei/notionapi"
)

var (
	client = notionapi.NewClient("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")
)

type NotionPageServiceImpl struct {
}

func (n *NotionPageServiceImpl) GetPageById(id string) *notionapi.Page {
	pageId := notionapi.PageID(id)
	page, err := client.Page.Get(context.Background(), pageId)
	if err != nil {
		if strings.Contains(err.Error(), "Could not find page with ID") {
			return nil
		}
		panic(err)
	}
	return page
}

func (n *NotionPageServiceImpl) CreateDatabasePage() *notionapi.Page {
	page, err := client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: "16626ea29ff180e3a9d1e83f92622638",
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: "New Page",
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
	return page
}
