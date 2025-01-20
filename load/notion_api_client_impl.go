package load

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

type NotionApiClient struct {
	client     *notionapi.Client
	databaseId string
}

func NewNotionApiClient(apiKey string, databaseId string) NotionApiClienter {
	token := notionapi.Token(apiKey)
	client := notionapi.NewClient(token)

	return &NotionApiClient{
		client:     client,
		databaseId: databaseId,
	}
}

func (n *NotionApiClient) CreateDatabasePage(properties []Property) string {
	notionApiProperties := notionapi.Properties{}

	for _, property := range properties {
		switch property := property.(type) {
		case TextProperty:
			switch property.GetName() {
			case "Name":
				notionApiProperties["Name"] = notionapi.TitleProperty{
					Title: []notionapi.RichText{
						{
							Text: &notionapi.Text{
								Content: property.Value,
							},
						},
					},
				}
			default:
				notionApiProperties[property.GetName()] = notionapi.RichTextProperty{
					RichText: []notionapi.RichText{
						{
							Text: &notionapi.Text{
								Content: property.Value,
							},
						},
					},
				}
			}
		default:
			msg := fmt.Sprint("Unsupported property type with name: ", property.GetName())
			panic(msg)
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

func (n NotionApiClient) QueryDatabasePageExists(property TextProperty) bool {
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
