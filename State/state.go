package State

import "fmt"

type Week interface {
	Today()
	Next(*DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

type Sunday struct{}

func (s *Sunday) Today() {
	fmt.Println("Sunday")
}

func (s *Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
	return
}

type Monday struct{}

func (s *Monday) Today() {
	fmt.Println("Monday")
}

func (s *Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
	return
}

type Tuesday struct{}

func (s *Tuesday) Today() {
	fmt.Println("Tuesday")
}

func (s *Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
	return
}

type Wednesday struct{}

func (s *Wednesday) Today() {
	fmt.Println("Wednesday")
}

func (s *Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
	return
}

type Thursday struct{}

func (s *Thursday) Today() {
	fmt.Println("Thursday")
}

func (s *Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
	return
}

type Friday struct{}

func (s *Friday) Today() {
	fmt.Println("Friday")
}

func (s *Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
	return
}

type Saturday struct{}

func (s *Saturday) Today() {
	fmt.Println("Saturday")
}

func (s *Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
	return
}
