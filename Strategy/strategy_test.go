package Strategy

import "testing"

func TestNewPayment(t *testing.T) {
	c := NewPayContext("mim", "123", 1000000000)
	s := NewWXPayment()
	z := NewZFBPayment()
	p := NewPayment(c, s)
	p.Pay()
	pz := NewPayment(c, z)
	pz.Pay()
}
