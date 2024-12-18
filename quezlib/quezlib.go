package quezlib

import "github.com/pluckynumbat/linked-list-stuff-go/tlistlib"

type Queue struct {
	list *tlistlib.TailedList
}

// Method to check whether a pointer to a Queue is nil
func (q *Queue) IsNil() bool {
	return q == nil
}

