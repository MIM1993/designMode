package Obeserver

import "testing"

func TestSubject(t *testing.T) {
	sub := NewSubject("hello everybody")
	r1 := NewReader("r1")
	r2 := NewReader("r2")
	r3 := NewReader("r3")

	sub.AppendObeserve(r2, r1, r3)
	sub.UpdateContext("")

	sub.UpdateContext("my name is OBM")
}
