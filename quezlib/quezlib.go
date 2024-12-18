package quezlib

import "github.com/pluckynumbat/linked-list-stuff-go/tlistlib"

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
