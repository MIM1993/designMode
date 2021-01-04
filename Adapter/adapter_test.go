package Adapter

import "testing"

var expect = "adaptee method"
var expectCopy = "adapteeImplCopy method"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}

	adapteeCp := NewAdapteeImplCopy()
	tg := NewAdapterCopy(adapteeCp)
	resCp := tg.Request()
	if resCp != expectCopy {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}

