package load

import "fmt"

func Upsert(id string) {
	page := GetPageById(id)
	if page == nil {
		fmt.Println("Page not found, creating new page")
		page = CreateDatabasePage()
	} else {
		fmt.Println("Page found")
	}
}
