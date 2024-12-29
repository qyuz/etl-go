package main

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func main() {

	client := notionapi.NewClient("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")

	// page, err := client.Database.Get(context.Background(), "16626ea29ff180e3a9d1e83f92622638")
	page, err := client.Page.Get(context.Background(), "16626ea29ff180f39493cb1425c8ba54")
	if err != nil {
		fmt.Println(err)
	}

	propertyName := "title"

	// Access the title property
	titleProperty, ok := page.Properties[propertyName].(*notionapi.ButtonProperty)
	// titleProperty, ok := page.Properties[propertyName].(*notionapi.Property)
	// titleProperty, ok := page.Properties[propertyName].(*notionapi.TitleProperty)
	if !ok {
		fmt.Println("Property is not a title property")
		return
	}
	fmt.Println("Property is a title property", titleProperty)
	fmt.Page
	// if len(titleProperty.Title) > 0 {
	// 	text := titleProperty.Title[0].Text.Content
	// 	fmt.Println("Title:", text)
	// } else {
	// 	fmt.Println("Title property is empty")
	// }

	fmt.Println(page.Properties)
	fmt.Println(page.Properties[propertyName])
	fmt.Println(page.Properties[propertyName].GetType())
}
