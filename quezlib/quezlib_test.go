package quezlib

import (
	"fmt"
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
	var q1, q2, q3, q4 *Queue
	q2 = &Queue{}
	q3 = &Queue{&tlistlib.TailedList{}}

	tl := &tlistlib.TailedList{}
	tl.AddAtEnd("a")
	q4 = &Queue{tl}

	var tests = []struct {
		name string
		q    *Queue
		want bool
	}{
		{"nil queue", q1, true},
		{"non nil queue, nil list", q2, true},
		{"non nil queue, empty list", q3, true},
		{"non nil queue, non empty list", q4, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.q.IsEmpty()
			if got != test.want {
				t.Errorf("isListNil gave incorrect results, want: %v, got %v", test.want, got)
			}
		})
	}
}

func TestPeekQueueTillEmpty(t *testing.T) {
	tl := &tlistlib.TailedList{}
	q := &Queue{tl}

	tl.AddAtEnd("a")
	tl.AddAtEnd("b")
	tl.AddAtEnd("c")

	want := "a"
	got, err := q.Peek()
	if err != nil {
		t.Errorf("Peek on the Queue failed with error: %v", err)
	}
	if want != got {
		t.Errorf("Peek gave incorrect results, want: %v, got %v", want, got)
	}

	_, err = tl.RemoveFirst()
	if err != nil {
		t.Errorf("RemoveFirst on list failed with error: %v", err)
	}

	want = "b"
	got, err = q.Peek()
	if err != nil {
		t.Errorf("Peek on the Queue failed with error: %v", err)
	}
	if want != got {
		t.Errorf("Peek gave incorrect results, want: %v, got %v", want, got)
	}

	_, err = tl.RemoveFirst()
	if err != nil {
		t.Errorf("RemoveFirst on list failed with error: %v", err)
	}

	want = "c"
	got, err = q.Peek()
	if err != nil {
		t.Errorf("Peek on the Queue failed with error: %v", err)
	}
	if want != got {
		t.Errorf("Peek gave incorrect results, want: %v, got %v", want, got)
	}

	_, err = tl.RemoveFirst()
	if err != nil {
		t.Errorf("RemoveFirst on list failed with error: %v", err)
	}

	_, err = q.Peek()
	if err != nil {
		fmt.Println(err)
	} else {
		t.Errorf("calling Peek() on an empty Queue should return an error: %v", err)
	}
}

func TestPeekNilOrEmptyQueue(t *testing.T) {
	var q1, q2, q3 *Queue
	q2 = &Queue{}
	q3 = &Queue{&tlistlib.TailedList{}}

	var tests = []struct {
		name string
		q    *Queue
		want error
	}{
		{"nil queue", q1, queueNilError},
		{"non nil queue, nil list", q2, queueEmptyError},
		{"non nil queue, empty list", q3, queueEmptyError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, got := test.q.Peek()
			if got != test.want {
				t.Errorf("Peek gave incorrect error, want: %v, got %v", test.want, got)
			}
		})
	}
}
