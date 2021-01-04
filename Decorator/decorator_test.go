package Decorator

import (
	"fmt"
	"testing"
)

func TestNewConcreteComponent(t *testing.T) {
	c := NewConcreteComponent(10)
	c = WarpMulDecorator(c, 10)
	fmt.Println(c.Output())
}


func TestNewConcreteComponent1(t *testing.T) {
	c := NewConcreteComponent(10)
	c = WarpAddDecorator(c, 10)
	fmt.Println(c.Output())
}
