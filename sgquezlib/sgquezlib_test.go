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
