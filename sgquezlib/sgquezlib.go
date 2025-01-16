// Package sgquezlib: library for a Semi generic queue that can contain data of any type that implements the fmt.Stringer interface
package sgquezlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"
)

type SemiGenericQueue[T fmt.Stringer] struct {
	sdlist *sdlistlib.SemiGenericList[T]
}
