package FactoryMethod

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	var p = PlusOperatorFactory{}
	cp := p.Create()
	cp.SetA(10)
	cp.SetB(20)
	fmt.Println(cp.Result())

	var m = MinusOperatorFactory{}
	cm := m.Create()
	cm.SetA(10)
	cm.SetB(20)
	fmt.Println(cm.Result())
}
