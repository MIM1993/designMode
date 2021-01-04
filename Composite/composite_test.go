package Composite

import "testing"

func TestComponent(t *testing.T) {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild([]Component{c1, c2})
	c1.AddChild([]Component{c3})
	c1.AddChild([]Component{l1})
	c2.AddChild([]Component{l2})
	c3.AddChild([]Component{l3})

	root.Print(" ")
}
