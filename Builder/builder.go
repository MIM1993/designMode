package Builder

type IBuilder interface {
	Part1()
	Part2()
	Part3()
	//ReturnResult() string
}

type Director struct {
	builder IBuilder
}

func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

//func (d *Director) GetResult() string {
//	return d.builder.ReturnResult()
//}

func NewDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "1"
}

func (b *Builder1) ReturnResult() string {
	return b.result
}
