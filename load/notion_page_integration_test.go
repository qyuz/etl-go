package load

import (
	"testing"
)

func TestGetPageById(t *testing.T) {
	notionPage := &NotionPageService{}
	page := notionPage.GetPageById("16626ea29ff180f39493cb1425c8ba54")
	if page == nil {
		t.Error("Page should not be nil")
	}
}

func TestGetPageByIdDoesNotExist(t *testing.T) {
	notionPage := &NotionPageService{}
	page := notionPage.GetPageById("16626ea29ff180f39493cb1425c8ba55")
	if page != nil {
		t.Error("Page should be nil")
	}
}
