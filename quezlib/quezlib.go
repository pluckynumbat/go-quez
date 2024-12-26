package quezlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/tlistlib"
)

var queueNilError = fmt.Errorf("The queue is nil")
var queueEmptyError = fmt.Errorf("The queue is empty")

type Queue struct {
	tlist *tlistlib.TailedList
}

// Function to create a new Queue with a new underlying list
func CreateNewQueue() *Queue {
	tl := &tlistlib.TailedList{}
	q := &Queue{tl}
	return q
}

// Method to check whether a pointer to a Queue is nil
func (q *Queue) IsNil() bool {
	return q == nil
}

// Internal method to check whether the underlying list is nil
func (q *Queue) isListNil() bool {
	return q.IsNil() || q.tlist.IsNil()
}

// Method to check whether a Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.IsNil() || q.tlist.IsNil() || q.tlist.Head() == nil
}

// Method to get the data at the beginning of the Queue
func (q *Queue) Peek() (string, error) {
	if q.IsNil() {
		return "", queueNilError
	}

	if q.IsEmpty() {
		return "", queueEmptyError
	}

	return q.tlist.Head().String(), nil
}

// Method to add an element at the end of the Queue
func (q *Queue) Enqueue(val string) error {
	if q.IsNil() {
		return queueNilError
	}

	if q.isListNil() {
		q.tlist = &tlistlib.TailedList{}
	}

	q.tlist.AddAtEnd(val)
	return nil
}
