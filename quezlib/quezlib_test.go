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

	var tests = []struct {
		name string
		want string
	}{
		{"3 elements queue", "a"},
		{"2 elements queue", "b"},
		{"1 element queue", "c"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := q.Peek()
			if err != nil {
				t.Errorf("Peek on the Queue failed with error: %v", err)
			}
			if got != test.want {
				t.Errorf("Peek gave incorrect results, want: %v, got %v", test.want, got)
			}

			_, err = tl.RemoveFirst()
			if err != nil {
				t.Errorf("RemoveFirst on list failed with error: %v", err)
			}
		})
	}

	_, err := q.Peek()
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

func TestEnqueueNilQueue(t *testing.T) {
	var q *Queue
	err := q.Enqueue("a")
	if err == nil {
		t.Errorf("Enqueue() on a nil queue should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestEnqueueEmptyQueue(t *testing.T) {
	q := &Queue{}
	err := q.Enqueue("a")
	if err != nil {
		t.Errorf("Enqueue() failed with error: %v", err)
	} else {
		want := "a"
		got, err2 := q.Peek()
		if err2 != nil {
			t.Errorf("Peek() failed with error: %v", err2)
		} else {
			if got != want {
				t.Errorf("Enqueue() gave incorrect results, want: %v, got: %v", want, got)
			}
		}
	}
}

func TestPeekAfterEnqueue(t *testing.T) {
	q := &Queue{}
	err := q.Enqueue("a")
	if err != nil {
		t.Errorf("Enqueue() failed with error: %v", err)
	} else {
		want := "a"
		got, err2 := q.Peek()
		if err2 != nil {
			t.Errorf("Peek() failed with error: %v", err2)
		} else {
			if got != want {
				t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
			}
		}
	}

	err = q.Enqueue("b")
	if err != nil {
		t.Errorf("Enqueue() failed with error: %v", err)
	} else {
		want := "a"
		got, err2 := q.Peek()
		if err2 != nil {
			t.Errorf("Peek() failed with error: %v", err2)
		} else {
			if got != want {
				t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
			}
		}
	}
	err = q.Enqueue("c")
	if err != nil {
		t.Errorf("Enqueue() failed with error: %v", err)
	} else {
		want := "a"
		got, err2 := q.Peek()
		if err2 != nil {
			t.Errorf("Peek() failed with error: %v", err2)
		} else {
			if got != want {
				t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
			}
		}
	}
}
