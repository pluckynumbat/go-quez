package sgquezlib

import (
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
