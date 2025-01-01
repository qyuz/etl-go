package load

import (
	"testing"
)

func TestCheckPageExists(t *testing.T) {
	runIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	pageExists := notionPage.CheckPageExists("16626ea29ff180f39493cb1425c8ba54")
	if !pageExists {
		t.Error("Page should exist")
	}
}

func TestCheckPageDoesNotExist(t *testing.T) {
	runIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	pageExists := notionPage.CheckPageExists("16626ea29ff180f39493cb1425c8ba55")
	if pageExists {
		t.Error("Page should not exist")
	}
}

func TestGetPageById(t *testing.T) {
	runIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	page := notionPage.GetPageById("16626ea29ff180f39493cb1425c8ba54")
	if page == nil {
		t.Error("Page should not be nil")
	}
}

func TestGetPageByIdDoesNotExist(t *testing.T) {
	runIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	page := notionPage.GetPageById("16626ea29ff180f39493cb1425c8ba55")
	if page != nil {
		t.Error("Page should be nil")
	}
}
