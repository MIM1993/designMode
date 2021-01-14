package State

import (
	"testing"
	"time"
)

func TestDayContext(t *testing.T) {
	day := NewDayContext()
	i := 0
	for ; i < 10; i++ {
		day.Today()
		day.Next()
		time.Sleep(time.Second)
	}
}