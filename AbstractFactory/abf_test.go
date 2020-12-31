package AbstractFactory_test

import (
	"designMode/AbstractFactory"
	"testing"
)

func TestAbf(t *testing.T){
	var af = AbstractFactory.AFty{}
	one := af.CreateOne()
	two := af.CreateTwo()
	one.DoOne()
	two.DoTwo()
}