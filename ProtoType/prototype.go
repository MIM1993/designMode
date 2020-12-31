package prototype

type Cloneable interface {
	Clone() Cloneable
}

type IClass interface {
	Cloneable
	SetName(string)
	GetName()string
}

