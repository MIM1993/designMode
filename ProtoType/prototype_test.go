package prototype

import (
	"fmt"
	"testing"
)

type class struct {
	Name string
}

func (c *class) Clone() Cloneable {
	cc := *c
	return &cc
}

func (c *class) SetName(name string) {
	c.Name = name
}

func (c *class) GetName() string {
	return c.Name
}


func TestProto(t *testing.T){
	c := class{Name:"mim"}
	fmt.Println(c.GetName())
	cc := c.Clone()
	ccc:= cc.(*class)
	ccc.SetName("nbdl")
	fmt.Println(ccc.GetName())
}