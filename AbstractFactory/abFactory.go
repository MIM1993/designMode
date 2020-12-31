package AbstractFactory

import "fmt"

type IAbFactory interface {
	CreateOne() IOne
	CreateTwo() ITwo
}

type IOne interface {
	DoOne()
}

type ITwo interface {
	DoTwo()
}

type AFty struct {
	one IOne
	two ITwo
}

func (A *AFty) CreateOne() IOne {
	return &one{}
}

func (A *AFty) CreateTwo() ITwo {
	return &two{}
}

type one struct{}

func (o *one) DoOne() {
	fmt.Println("this is one class")
}

type two struct{}

func (t *two) DoTwo() {
	fmt.Println("this is two class")
}
