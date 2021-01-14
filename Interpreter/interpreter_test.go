package Interpreter

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	p :=  &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6 + 10")
	res := p.Result().interpret()
	fmt.Println(res)

	//sub&mode != 0
	fmt.Println(3&3)
}