package load

import (
	"context"
	"strings"

	"github.com/jomei/notionapi"
)

var (
	client     = notionapi.NewClient("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")
	databaseID = notionapi.DatabaseID("16626ea29ff180e3a9d1e83f92622638")
)

type NotionPageServiceImpl struct {
}

func (n *NotionPageServiceImpl) CheckDatabasePageExists(property string, id string) bool {
	DatabaseQueryResponse, err := client.Database.Query(context.Background(), databaseID, &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "Movie ID",
			RichText: &notionapi.TextFilterCondition{Equals: id},
		}})

	if err != nil {
		panic(err)
	}

	return len(DatabaseQueryResponse.Results) > 0
}

func (n *NotionPageServiceImpl) CreateDatabasePage(properties []Property) string {
	notionApiProperties := notionapi.Properties{}
	for _, property := range properties {
		switch property.GetName() {
		case "Name":
			notionApiProperties["Name"] = notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: property.(TextProperty).Text,
						},
					},
				},
			}
		default:
			notionApiProperties[property.GetName()] = notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: property.(TextProperty).Text,
						},
					},
				},
			}
		}
	}

	page, err := client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: "16626ea29ff180e3a9d1e83f92622638",
		},
		Properties: notionApiProperties,
	})
	if err != nil {
		panic(err)
	}
	return page.ID.String()
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
