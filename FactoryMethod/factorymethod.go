package FactoryMethod

type IOperater interface {
	SetA(int)
	SetB(int)
	Result() int
}

type IFactory interface {
	Create() IOperater
}

type BaseOperate struct {
	A int
	B int
}

func (base *BaseOperate) SetA(a int) {
	base.A = a
}

func (base *BaseOperate) SetB(b int) {
	base.B = b
}

type PlusOperatorFactory struct{}

type PlusOperator struct {
	*BaseOperate
}

func (p *PlusOperator) Result() int {
	return p.A + p.B
}

func (p *PlusOperatorFactory) Create() IOperater {
	return &PlusOperator{
		BaseOperate: &BaseOperate{},
	}
}

type MinusOperatorFactory struct{}

type MinusOperator struct {
	*BaseOperate
}

func (m *MinusOperator) Result() int {
	return m.A * m.B
}

func (m *MinusOperatorFactory) Create() IOperater {
	return &MinusOperator{
		BaseOperate: &BaseOperate{},
	}
}

