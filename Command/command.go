package Command

import "fmt"

type Command interface {
	Exec()
}

type motherBoard struct{}

func NewMotherBoard() *motherBoard {
	return &motherBoard{}
}

func (m *motherBoard) SystemStart() {
	fmt.Printf("invoke SystemStart\n")
}

func (m *motherBoard) SystemRobot() {
	fmt.Printf("invoke Robot\n")
}

type startCommand struct {
	mb *motherBoard
}

func NewStartCommand(mb *motherBoard) Command {
	return &startCommand{
		mb: mb,
	}
}

func (s *startCommand) Exec() {
	s.mb.SystemStart()
}

type rebotCommand struct {
	mb *motherBoard
}

func NewRebotCommand(mb *motherBoard) Command {
	return &rebotCommand{
		mb: mb,
	}
}

func (r *rebotCommand) Exec() {
	r.mb.SystemRobot()
}

type box struct {
	button1 Command
	button2 Command
}

func NewBox(button1, button2 Command) *box {
	return &box{
		button1: button1,
		button2: button2,
	}
}

func (b *box) PressButton1() {
	b.button1.Exec()
}

func (b *box) PressButton2() {
	b.button2.Exec()
}

type batchBox struct {
	buttons []Command
}

func NewBatchBox() *batchBox {
	return &batchBox{
		buttons: make([]Command, 0),
	}
}

func (b *batchBox) AppendCommand(com ...Command) {
	b.buttons = append(b.buttons, com...)
}

func (b *batchBox) Batch() {
	for _, v := range b.buttons {
		v.Exec()
	}
}
