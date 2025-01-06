package load

import (
	"context"

	"github.com/jomei/notionapi"
)

type NotionApiClientImpl struct {
	apiKey     string
	client     *notionapi.Client
	databaseId string
}

func (n *NotionApiClientImpl) CreateDatabasePage(properties []Property) string {
	notionApiProperties := notionapi.Properties{}

	for _, property := range properties {
		switch property.GetName() {
		case "Name":
			notionApiProperties["Name"] = notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: property.(TextProperty).Value,
						},
					},
				},
			}
		default:
			notionApiProperties[property.GetName()] = notionapi.RichTextProperty{
				RichText: []notionapi.RichText{
					{
						Text: &notionapi.Text{
							Content: property.(TextProperty).Value,
						},
					},
				},
			}
		}
	}

	page, err := n.client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			DatabaseID: notionapi.DatabaseID(n.databaseId),
		},
		Properties: notionApiProperties,
	})

	if err != nil {
		panic(err)
	}

	return page.ID.String()
}

func (n NotionApiClientImpl) QueryDatabasePageExists(property TextProperty) bool {
	DatabaseQueryResponse, err := n.client.Database.Query(context.Background(), notionapi.DatabaseID(n.databaseId), &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: property.GetName(),
			RichText: &notionapi.TextFilterCondition{Equals: property.Value},
		}})

	if err != nil {
		panic(err)
	}

	return len(DatabaseQueryResponse.Results) > 0
}
