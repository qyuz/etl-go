package load

import (
	"testing"
)

func TestCheckPageExists(t *testing.T) {
	verifyShouldRunIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	pageExists := notionPage.CheckDatabasePageExists("Movie ID", "Movie ID")
	if !pageExists {
		t.Error("Page should exist")
	}
}

func TestCheckPageDoesNotExist(t *testing.T) {
	verifyShouldRunIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	pageExists := notionPage.CheckDatabasePageExists("Movie ID", "Movie ID that does not exist")
	if pageExists {
		t.Error("Page should not exist")
	}
}

func TestGetPageById(t *testing.T) {
	verifyShouldRunIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	page := notionPage.GetPageById("16626ea29ff180f39493cb1425c8ba54")
	if page == nil {
		t.Error("Page should not be nil")
	}
}

func TestGetPageByIdDoesNotExist(t *testing.T) {
	verifyShouldRunIntegrationTests(t)
	notionPage := &NotionPageServiceImpl{}
	page := notionPage.GetPageById("16626ea29ff180f39493cb1425c8ba55")
	if page != nil {
		t.Error("Page should be nil")
	}
}
