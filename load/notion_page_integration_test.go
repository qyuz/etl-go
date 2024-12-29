package load

import (
	"testing"
)

func TestGetPageById(t *testing.T) {
	page := GetPageById("16626ea29ff180f39493cb1425c8ba54")
	if page == nil {
		t.Error("Page should not be nil")
	}
}

func TestGetPageByIdDoesNotExist(t *testing.T) {
	page := GetPageById("16626ea29ff180f39493cb1425c8ba55")
	if page != nil {
		t.Error("Page should be nil")
	}
}

func TestCreateDatabasePage(t *testing.T) {
	page := CreateDatabasePage()
	if page == nil {
		t.Error("Page should not be nil")
	}
	pageById := GetPageById(page.ID.String())
	if pageById == nil {
		t.Error("Should be able to get page by ID")
	}
}
