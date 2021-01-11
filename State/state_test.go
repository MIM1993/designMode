package State

import (
	"testing"
	"time"
)

func TestDayContext(t *testing.T) {
	day := NewDayContext()
	for {
		day.Today()
		day.Next()
	 	time.Sleep(time.Second)
	}
}