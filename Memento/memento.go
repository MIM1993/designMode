package Memento

import "fmt"

type Memento interface {
	reduction(*Game)
}

type Game struct {
	hp, mp int
}

func NewGame(hp, mp int) *Game {
	return &Game{
		hp: hp,
		mp: mp,
	}
}

type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(hp, mp int) {
	g.hp += hp
	g.mp += mp
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(me Memento) {
	me.reduction(g)
}

func (g *gameMemento) reduction(game *Game) {
	game.hp = g.hp
	game.mp = g.mp
}

func (g *Game) Status() {
	fmt.Printf("hp:%d , mp:%d\n", g.hp, g.mp)
}
