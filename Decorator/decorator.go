package Decorator

type Component interface {
	Output() int
}

type ConcreteComponent struct {
	num int
}

func NewConcreteComponent(num int) Component {
	return &ConcreteComponent{
		num: num,
	}
}

func (c *ConcreteComponent) Output() int {
	return c.num
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (c *MulDecorator) Output() int {
	return c.num * c.Component.Output()
}

type AddDecorator struct {
	Component
	num int
}

func (a *AddDecorator) Output() int {
	return a.num + a.Component.Output()
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}
