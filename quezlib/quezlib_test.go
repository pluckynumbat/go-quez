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
		t.Error("calling Peek() on an empty Queue should return an error")
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
		t.Error("Enqueue() on a nil queue should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestPeekAfterEnqueue(t *testing.T) {
	q := &Queue{}

	var tests = []struct {
		name string
		val  string
	}{
		{"1 element queue", "a"},
		{"2 element queue", "b"},
		{"3 element queue", "c"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := q.Enqueue(test.val)
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
		})
	}
}

func TestDequeueNilQueue(t *testing.T) {
	var q *Queue

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Dequeue() on a nil Queue should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := &Queue{}

	_, err := q.Dequeue()
	if err == nil {
		t.Error("Dequeue() on an empty Queue should have returned an error")
	} else {
		fmt.Println(err)
	}
}

func TestDequeueTillEmpty(t *testing.T) {
	tl := &tlistlib.TailedList{}
	tl.AddAtEnd("a")
	tl.AddAtEnd("b")
	tl.AddAtEnd("c")

	q := &Queue{tl}

	var tests = []struct {
		name       string
		dequeueVal string
		newPeek    string
		expPeekErr error
	}{
		{"3 elements", "a", "b", nil},
		{"2 elements", "b", "c", nil},
		{"1 element", "c", "", queueEmptyError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := q.Dequeue()
			if err != nil {
				t.Errorf("Dequeue() failed with error: %v", err)
			} else {
				want := test.dequeueVal
				got := val
				if want != got {
					t.Errorf("Dequeue() gave incorrect results, want: %v, got: %v", want, got)
				}
			}

			val2, err2 := q.Peek()
			if err2 != test.expPeekErr {
				t.Errorf("Peek() error doesn't match expected error, want: %v, got: %v", test.expPeekErr, err2)
			} else {
				want := test.newPeek
				got := val2
				if want != got {
					t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
				}
			}
		})
	}
}

func TestQueueOperations(t *testing.T) {
	q := &Queue{}

	var enqueueTests = []struct {
		name string
		val  string
	}{
		{"1 element queue", "a"},
		{"2 element queue", "b"},
		{"3 element queue", "c"},
	}

	for _, test := range enqueueTests {
		t.Run(test.name, func(t *testing.T) {
			err := q.Enqueue(test.val)
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
		})
	}

	var dequeueTests = []struct {
		name       string
		dequeueVal string
		newPeek    string
		expPeekErr error
	}{
		{"3 elements", "a", "b", nil},
		{"2 elements", "b", "c", nil},
		{"1 element", "c", "", queueEmptyError},
	}

	for _, test := range dequeueTests {
		t.Run(test.name, func(t *testing.T) {
			val, err := q.Dequeue()
			if err != nil {
				t.Errorf("Dequeue() failed with error: %v", err)
			} else {
				want := test.dequeueVal
				got := val
				if want != got {
					t.Errorf("Dequeue() gave incorrect results, want: %v, got: %v", want, got)
				}
			}

			val2, err2 := q.Peek()
			if err2 != test.expPeekErr {
				t.Errorf("Peek() error doesn't match expected error, want: %v, got: %v", test.expPeekErr, err2)
			} else {
				want := test.newPeek
				got := val2
				if want != got {
					t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
				}
			}
		})
	}

	var stateTests = []struct {
		name   string
		fn     func() bool
		fnName string
		want   bool
	}{
		{"is nil", q.IsNil, "IsNil", false},
		{"is list nil", q.isListNil, "isListNil", false},
		{"is empty", q.IsEmpty, "isEmpty", true},
	}

	for _, test := range stateTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.fn()
			if got != test.want {
				t.Errorf("Got incorrect results for the state function %v, want: %v, got: %v", test.fnName, test.want, got)
			}
		})
	}
}
