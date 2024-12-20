package quezlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/tlistlib"
)

var queueNilError = fmt.Errorf("The queue is nil")
var queueEmptyError = fmt.Errorf("The queue is empty")

type Queue struct {
	list *tlistlib.TailedList
}

// Method to check whether a pointer to a Queue is nil
func (q *Queue) IsNil() bool {
	return q == nil
}

// Internal method to check whether the underlying list is nil
func (q *Queue) isListNil() bool {
	return q.IsNil() || q.list.IsNil()
}

// Method to check whether a Queue is empty
func (q *Queue) IsEmpty() bool {
	return q.IsNil() || q.list.IsNil() || q.list.Head() == nil
}
