package load

import (
	"context"
	"strings"

	"github.com/jomei/notionapi"
)

var (
	client = notionapi.NewClient("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")
)

func getPageById(id string) *notionapi.Page {
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
