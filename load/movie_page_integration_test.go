package load

import "testing"

func TestUpsertPageFound(t *testing.T) {
	moviePage := &MoviePage{ID: "16626ea29ff180f39493cb1425c8ba54"}
	moviePage.Upsert()
}

func TestUpsertPageNotFound(t *testing.T) {
	moviePage := &MoviePage{ID: "15d26ea29ff180d1b838ed84816e867d"}
	moviePage.Upsert()
}
