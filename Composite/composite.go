package Composite

import (
	"fmt"
	"log"
)

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild([]Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(flag int, name string) Component {
	var c Component
	switch flag {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}
	if c == nil {
		log.Fatal("NewComponent error")
		return nil
	}
	c.SetName(name)
	return c
}

//base class
type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(p Component) {
	c.parent = p
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild([]Component) {}

func (c *component) Print(string) {}

//leaf class
type Leaf struct {
	component
}

func NewLeaf() Component {
	return &Leaf{}
}

func (l *Leaf) Print(str string) {
	fmt.Printf("%s-%s\n", str, l.name)
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() Component {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) AddChild(childs []Component) {
	for _, child := range childs {
		child.SetParent(c)
	}
	c.childs = append(c.childs, childs...)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.name)
	pre += " "
	for _, child := range c.childs {
		child.Print(pre)
	}
}
