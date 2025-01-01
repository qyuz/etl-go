package load

import (
	"testing"
)

type NotionPageService interface {
	CheckPageExists(id string) bool
	CreateDatabasePage(id string) string
}

func runIntegrationTests(t *testing.T) {
	if false {
		t.Skip("skipping integration test")
	}
}
