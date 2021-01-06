package Iterator

import (
	"fmt"
	"testing"
)

func TestNumbersIterator(t *testing.T) {
	var ag Aggregate
	ag = NewNumbers(1, 10)
	numIter := ag.Iterator()

	for numIter.First(); !numIter.IsDone(); {
		fmt.Println(numIter.Next())
	}
}
