package quezlib

import (
	"testing"

	"github.com/pluckynumbat/linked-list-stuff-go/tlistlib"
)

func TestIsNil(t *testing.T) {
	var q1, q2 *Queue
	q2 = &Queue{}

	var tests = []struct {
		name string
		q    *Queue
		want bool
	}{
		{"nil true", q1, true},
		{"nil false", q2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.q.IsNil()
			if got != test.want {
				t.Errorf("IsNil gave incorrect results, want: %v, got %v", test.want, got)
			}
		})
	}
}

func TestIsListNil(t *testing.T) {
	var q1, q2, q3 *Queue
	q2 = &Queue{}
	q3 = &Queue{&tlistlib.TailedList{}}

	var tests = []struct {
		name string
		q    *Queue
		want bool
	}{
		{"nil queue", q1, true},
		{"nil list", q2, true},
		{"nil false", q3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.q.isListNil()
			if got != test.want {
				t.Errorf("isListNil gave incorrect results, want: %v, got %v", test.want, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var q *Queue

	want := true
	got := q.IsEmpty()

	if want != got {
		t.Errorf("IsEmpty gave incorrect results, want: %v, got %v", want, got)
	}
}

