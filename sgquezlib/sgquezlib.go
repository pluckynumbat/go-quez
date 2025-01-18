// Package sgquezlib: library for a Semi generic queue that can contain data of any type that implements the fmt.Stringer interface
package sgquezlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"
)

var queueNilError = fmt.Errorf("the queue is nil")
var queueEmptyError = fmt.Errorf("the queue is empty")

type SemiGenericQueue[T fmt.Stringer] struct {
	sdlist *sdlistlib.SemiGenericList[T]
}

// IsNil checks whether a pointer to a Semi Generic Queue is nil
func (queue *SemiGenericQueue[T]) IsNil() bool {
	return queue == nil
}

// Internal Method to check whether the underlying list is nil
func (queue *SemiGenericQueue[T]) isListNil() bool {
	return queue.IsNil() || queue.sdlist.IsNil()
}

// IsEmpty checks whether a Semi Generic Queue is empty
func (queue *SemiGenericQueue[T]) IsEmpty() bool {
	return queue.IsNil() || queue.isListNil() || queue.sdlist.IsEmpty()
}
