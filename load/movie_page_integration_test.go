package load

import "testing"

func skipMoviePageIntegrationTests(t *testing.T) {
	if true {
		t.Skip("skipping test")
	}
}

// replace skipTests with skipMoviePageTests

func TestUpsertPageFound(t *testing.T) {
	skipMoviePageIntegrationTests(t)
	moviePage := &MoviePage{ID: "16626ea29ff180f39493cb1425c8ba54", notionPageService: &NotionPageServiceImpl{}}
	moviePage.Upsert()
}

func TestUpsertPageNotFound(t *testing.T) {
	skipMoviePageIntegrationTests(t)
	moviePage := &MoviePage{ID: "15d26ea29ff180d1b838ed84816e867d", notionPageService: &NotionPageServiceImpl{}}
	moviePage.Upsert()
}
