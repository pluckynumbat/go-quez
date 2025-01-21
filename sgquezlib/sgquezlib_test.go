package sgquezlib

import (
	"errors"
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"

	"testing"
)

type prInt int       // printable int
type prString string // printable string

func (p prInt) String() string {
	return fmt.Sprintf("%v", int(p))
}

func (p prString) String() string {
	return fmt.Sprintf("%v", string(p))
}

func TestIsNil(t *testing.T) {

	var q1 *SemiGenericQueue[*prInt]
	q2 := &SemiGenericQueue[*prInt]{}

	var tests = []struct {
		name  string
		queue *SemiGenericQueue[*prInt]
		want  bool
	}{
		{"nil true", q1, true},
		{"nil false", q2, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.queue.IsNil()
			want := test.want

			if got != want {
				t.Errorf("IsNil() gave incorrect results, want: %v, got : %v", want, got)
			}
		})
	}
}

func TestIsListNil(t *testing.T) {

	var q1 *SemiGenericQueue[*prInt]
	q2 := &SemiGenericQueue[*prInt]{}

	l := &sdlistlib.SemiGenericList[*prInt]{}
	q3 := &SemiGenericQueue[*prInt]{l}

	var tests = []struct {
		name  string
		queue *SemiGenericQueue[*prInt]
		want  bool
	}{
		{"nil queue", q1, true},
		{"nil list", q2, true},
		{"nil queue", q3, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.queue.isListNil()
			want := test.want

			if got != want {
				t.Errorf("isListNil() gave incorrect results, want: %v, got : %v", want, got)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {

	var q1 *SemiGenericQueue[*prString]
	q2 := &SemiGenericQueue[*prString]{}

	l := &sdlistlib.SemiGenericList[*prString]{}
	q3 := &SemiGenericQueue[*prString]{l}

	var pr prString = "a"
	var ptr = &pr
	l2 := &sdlistlib.SemiGenericList[*prString]{}
	addErr := l2.AddAtBeginning(ptr)
	if addErr != nil {
		t.Fatalf("AddAtBeginning() failed with error: %v", addErr)
	}
	q4 := &SemiGenericQueue[*prString]{l2}

	var tests = []struct {
		name  string
		queue *SemiGenericQueue[*prString]
		want  bool
	}{
		{"nil queue", q1, true},
		{"non nil queue, nil list", q2, true},
		{"non nil queue, empty list", q3, true},
		{"non nil queue, non empty list", q4, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.queue.IsEmpty()
			want := test.want

			if got != want {
				t.Errorf("IsEmpty() gave incorrect results, want: %v, got : %v", want, got)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	var q1, q2, q3, q4 *SemiGenericQueue[*prInt]

	q2 = &SemiGenericQueue[*prInt]{}

	l1 := &sdlistlib.SemiGenericList[*prInt]{}
	q3 = &SemiGenericQueue[*prInt]{l1}

	l2 := &sdlistlib.SemiGenericList[*prInt]{}
	var pr1 prInt
	addErr := l2.AddAtBeginning(&pr1)
	if addErr != nil {
		t.Fatalf("list add failed with error: %v", addErr)
	}
	q4 = &SemiGenericQueue[*prInt]{l2}

	var tests1 = []struct {
		name     string
		queue    *SemiGenericQueue[*prInt]
		wantVal  *prInt
		expError error
	}{
		{"nil queue", q1, nil, queueNilError},
		{"non-nil queue, nil list", q2, nil, queueEmptyError},
		{"empty queue", q3, nil, queueEmptyError},
		{"non-empty queue", q4, &pr1, nil},
	}

	for _, test := range tests1 {
		t.Run(test.name, func(t *testing.T) {
			gotVal, gotErr := test.queue.Peek()
			if !errors.Is(gotErr, test.expError) {
				t.Errorf("Unexpected error for Peek(), want: %v, got : %v", test.expError, gotErr)
			} else if gotErr != nil {
				fmt.Println(gotErr)
			} else if gotVal != test.wantVal {
				t.Errorf("Incorrect result for Peek(), want: %v, got : %v", test.wantVal, gotVal)
			}
		})
	}

	l := &sdlistlib.SemiGenericList[prString]{}
	prStrs := []prString{"a", "b", "c"}

	for _, prStr := range prStrs {
		addErr := l.AddAtEnd(prStr)
		if addErr != nil {
			t.Fatalf("list add failed with error: %v", addErr)
		}
	}

	q := &SemiGenericQueue[prString]{l}

	var tests2 = []struct {
		name    string
		wantVal prString
	}{
		{"3 element queue", "a"},
		{"2 element queue", "b"},
		{"1 element queue", "c"},
	}

	for _, test := range tests2 {
		t.Run(test.name, func(t *testing.T) {
			gotVal, err := q.Peek()
			if err != nil {
				t.Fatalf("Peek() failed with error: %v", err)
			} else if gotVal != test.wantVal {
				t.Errorf("Incorrect result for Peek(), want: %v, got : %v", test.wantVal, gotVal)
			}

			_, remErr := l.RemoveFirst()
			if remErr != nil {
				t.Fatalf("list's RemoveFirst() failed with error: %v", remErr)
			}
		})
	}
}

func TestEnqueue(t *testing.T) {

	t.Run("nil queue of prString pointers", func(t *testing.T) {
		var q *SemiGenericQueue[*prString]
		var s prString = "a"
		err := q.Enqueue(&s)

		if err == nil {
			t.Errorf("Enqueue() on a nil queue should have returned an error")
		} else {
			fmt.Println(err)
		}
	})

	t.Run("non nil queue of prString pointers", func(t *testing.T) {
		q := &SemiGenericQueue[*prString]{}

		var tests = []struct {
			name string
			val  prString
		}{
			{"1 element queue", "a"},
			{"2 element queue", "b"},
			{"3 element queue", "c"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := q.Enqueue(&test.val)
				if err != nil {
					t.Errorf("Enqueue() failed with error: %v", err)
				} else {
					val, pErr := q.Peek()
					if pErr != nil {
						t.Errorf("Peek() failed with error: %v", pErr)
					} else {
						var want prString = "a"
						got := *val
						if got != want {
							t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
						}
					}
				}
			})
		}
	})

	t.Run("nil queue of prInt values", func(t *testing.T) {
		var q *SemiGenericQueue[prInt]
		var i prInt = 1
		err := q.Enqueue(i)

		if err == nil {
			t.Errorf("Enqueue() on a nil queue should have returned an error")
		} else {
			fmt.Println(err)
		}
	})

	t.Run("non nil queue of prInt values", func(t *testing.T) {
		q := &SemiGenericQueue[prInt]{}

		var tests = []struct {
			name string
			val  prInt
		}{
			{"1 element queue", 1},
			{"2 element queue", 2},
			{"3 element queue", 3},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				err := q.Enqueue(test.val)
				if err != nil {
					t.Errorf("Enqueue() failed with error: %v", err)
				} else {
					val, pErr := q.Peek()
					if pErr != nil {
						t.Errorf("Peek() failed with error: %v", pErr)
					} else {
						var want prInt = 1
						got := val
						if got != want {
							t.Errorf("Peek() gave incorrect results, want: %v, got: %v", want, got)
						}
					}
				}
			})
		}
	})
}

func TestDequeueNilEmptyQueue(t *testing.T) {

	t.Run("nil / empty queue of prStrings", func(t *testing.T) {
		var q1, q2, q3, q4 *SemiGenericQueue[prString]

		q2 = &SemiGenericQueue[prString]{}

		l1 := &sdlistlib.SemiGenericList[prString]{}
		q3 = &SemiGenericQueue[prString]{l1}

		l2 := &sdlistlib.SemiGenericList[prString]{}
		addErr := l2.AddAtEnd("a")
		if addErr != nil {
			t.Fatalf("Error while adding to a semi generic list: %v", addErr)
		}
		q4 = &SemiGenericQueue[prString]{l2}

		var tests = []struct {
			name     string
			queue    *SemiGenericQueue[prString]
			expPeek  prString
			expError error
		}{
			{"nil queue", q1, "", queueNilError},
			{"non-nil queue, nil list", q2, "", queueEmptyError},
			{"empty queue", q3, "", queueEmptyError},
			{"non-empty queue ", q4, "a", nil},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				val, err := test.queue.Dequeue()
				if !errors.Is(err, test.expError) {
					t.Errorf("Dequeue() failed with unexpected error: %v", err)
				} else if err != nil {
					fmt.Println(err)
				} else {
					want := test.expPeek
					got := val
					if got != want {
						t.Errorf("Dequeue() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			})
		}
	})

	t.Run("nil / empty queue of pointers to prInts", func(t *testing.T) {
		var q1, q2, q3, q4 *SemiGenericQueue[*prInt]

		q2 = &SemiGenericQueue[*prInt]{}

		l1 := &sdlistlib.SemiGenericList[*prInt]{}
		q3 = &SemiGenericQueue[*prInt]{l1}

		l2 := &sdlistlib.SemiGenericList[*prInt]{}
		var pr prInt = 1
		addErr := l2.AddAtEnd(&pr)
		if addErr != nil {
			t.Fatalf("Error while adding to a semi generic list: %v", addErr)
		}
		q4 = &SemiGenericQueue[*prInt]{l2}

		var tests = []struct {
			name     string
			queue    *SemiGenericQueue[*prInt]
			expPeek  *prInt
			expError error
		}{
			{"nil queue", q1, nil, queueNilError},
			{"non-nil queue, nil list", q2, nil, queueEmptyError},
			{"empty queue", q3, nil, queueEmptyError},
			{"non-empty queue ", q4, &pr, nil},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				val, err := test.queue.Dequeue()
				if !errors.Is(err, test.expError) {
					t.Errorf("Dequeue() failed with unexpected error: %v", err)
				} else if err != nil {
					fmt.Println(err)
				} else {
					want := test.expPeek
					got := val
					if got != want {
						t.Errorf("Dequeue() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			})
		}
	})
}

				}
			}
		})
	}
}

func TestDequeueTillQueueEmpty(t *testing.T) {
	q := &SemiGenericQueue[prInt]{}

	prInts := []prInt{1, 2, 3}

	for _, val := range prInts {
		addArr := q.Enqueue(val)
		if addArr != nil {
			t.Fatalf("Enqueue() failed with error: %v", addArr)
		}
	}

	var tests = []struct {
		name       string
		expVal     prInt
		newPeek    prInt
		expPeekErr error
	}{
		{"3 element queue", 1, 2, nil},
		{"2 element queue", 2, 3, nil},
		{"1 element queue", 3, 0, queueEmptyError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			val, err := q.Dequeue()
			if err != nil {
				t.Errorf("Dequeue() failed with error: %v", err)
			} else {
				if val != test.expVal {
					t.Errorf("Dequeue() returned incorrect results, want: %v, got: %v", test.expVal, val)
				}

				val2, err2 := q.Peek()
				if !errors.Is(err2, test.expPeekErr) {
					t.Errorf("Peek() failed with unexpected error: %v", err2)
				} else if err2 != nil {
					fmt.Println(err2)
				} else {
					want := test.newPeek
					got := val2
					if got != want {
						t.Errorf("Peek() returned incorrect results, want: %v, got: %v", want, got)
					}
				}
			}
		})
	}
}
