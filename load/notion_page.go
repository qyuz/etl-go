package load

import (
	"testing"
)

type NotionPageService interface {
	CheckDatabasePageExists(property string, id string) bool
	CreateDatabasePage(id string) string
}

func verifyShouldRunIntegrationTests(t *testing.T) {
	if false {
		t.Skip("skipping integration test")
	}
}
