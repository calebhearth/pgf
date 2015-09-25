package pgp

import "testing"

func TestSetAdd(t *testing.T) {
	got := set{}
	got.Add("string")

	if len(got) != 1 {
		t.Fatalf("Should only have 1 element, had %d", len(got))
	}

	if _, ok := got["string"]; !ok {
		t.Fatal(`Should have key "string"`)
	}

	got.Add("another string")

	if len(got) != 2 {
		t.Fatal("Adding a different string should add a new element")
	}

	if _, ok := got["another string"]; !ok {
		t.Fatal(`Should have key "another string"`)
	}

	got.Add("string")

	if len(got) != 2 {
		t.Fatal(`Should not add a second instance of "string"`)
	}
}

func TestSetReduce(t *testing.T) {
	s := set(map[string]struct{}{
		"one": struct{}{},
		"two": struct{}{},
	})

	want := []string{"oneone", "twotwo"}

	got := s.Reduce(func(v string) string {
		return v + v
	})

	if len(got) != len(want) {
		t.Fatal("Should be same length")
	}

	for i := range got {
		if got[i] != want[i] {
			t.Fatal("Wrong results")
		}
	}
}
