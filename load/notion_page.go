package load

import (
	"testing"
)

type NotionPageService interface {
	CheckDatabasePageExists(property string, id string) bool
	CreateDatabasePage(properties []Property) string
}

type Property interface {
	GetName() string
	GetType() string
}

type TextProperty struct {
	Name string
	Text string
}

func (p TextProperty) GetName() string {
	return p.Name
}

func (p TextProperty) GetType() string {
	return "text"
}

func verifyShouldRunIntegrationTests(t *testing.T) {
	if true {
		t.Skip("skipping integration test")
	}
}
