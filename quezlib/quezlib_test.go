package quezlib

import (
	"testing"
)

func TestIsNil(t *testing.T) {
	var q *Queue

	want := true
	got := q.IsNil()
	if got != want {
		t.Errorf("IsNil gave incorrect results, want: %v, got %v", want, got)
	}

	q = &Queue{}
	want = false
	got = q.IsNil()
	if got != want {
		t.Errorf("IsNil gave incorrect results, want: %v, got %v", want, got)
	}
}
