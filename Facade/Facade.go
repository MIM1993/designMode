package Facade

import "fmt"

type IAPI interface {
	Test()
}

type APIMode struct {
	A IAPIA
	B IAPIB
}

func (A APIMode) Test() {
	A.A.TestA()
	A.B.TestB()
	fmt.Println("Test...")
}

var _ IAPI = (*APIMode)(nil)

func NewIAPI() IAPI {
	return &APIMode{
		A: NewIAPIA(),
		B: NewIAPIB(),
	}
}

type IAPIA interface {
	TestA()
}

type APIModeA struct{}

func (A APIModeA) TestA() {
	fmt.Println("TestA...")
}

var _ IAPIA = (*APIModeA)(nil)

func NewIAPIA() IAPIA {
	return &APIModeA{}
}

type IAPIB interface {
	TestB()
}

type APIModeB struct{}

func (A APIModeB) TestB() {
	fmt.Println("TestB...")
}

var _ IAPIB = (*APIModeB)(nil)

func NewIAPIB() IAPIB {
	return &APIModeB{}
}
