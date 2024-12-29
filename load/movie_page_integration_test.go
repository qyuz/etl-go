package load

import "testing"

func TestUpsertPageFound(t *testing.T) {
	Upsert("16626ea29ff180f39493cb1425c8ba54")
}

func TestUpsertPageNotFound(t *testing.T) {
	Upsert("15d26ea29ff180d1b838ed84816e867d")
}
