package Memento

import "testing"

func TestGame(t *testing.T) {
	g := NewGame(10, 10)
	g.Status()
	mm := g.Save()
	g.Play(-2, -4)
	g.Status()
	g.Load(mm)
	g.Status()
}
