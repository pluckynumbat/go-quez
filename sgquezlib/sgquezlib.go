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

// Peek returns the data (if present) at the front of the queue
func (queue *SemiGenericQueue[T]) Peek() (T, error) {
	if queue.IsNil() {
		return *new(T), queueNilError
	}

	if queue.IsEmpty() {
		return *new(T), queueEmptyError
	}

	data, err := queue.sdlist.Head().GetData()
	if err != nil {
		return *new(T), fmt.Errorf("queue Peek() failed with error %v", err)
	}

	return data, nil
}

// Enqueue adds an element to the back of the queue
func (queue *SemiGenericQueue[T]) Enqueue(val T) error {
	if queue.IsNil() {
		return queueNilError
	}

	if queue.isListNil() {
		queue.sdlist = &sdlistlib.SemiGenericList[T]{}
	}

	err := queue.sdlist.AddAtEnd(val)
	if err != nil {
		return fmt.Errorf("queue Enqueue() failed with error %v", err)
	}
	return nil
}

// Dequeue removes and returns the element at the front of the queue
func (queue *SemiGenericQueue[T]) Dequeue() (T, error) {
	if queue.IsNil() {
		return *new(T), queueNilError
	}

	if queue.IsEmpty() {
		return *new(T), queueEmptyError
	}

	node, err := queue.sdlist.RemoveFirst()
	if err != nil {
		return *new(T), fmt.Errorf("queue Dequeue() failed with error %v", err)
	}

	data, err2 := node.GetData()
	if err2 != nil {
		return *new(T), fmt.Errorf("queue Dequeue() failed with error %v", err)
	}

	return data, nil
}
