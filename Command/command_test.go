package Command_test

import (
	"designMode/Command"
	"testing"
)

func TestCommand(t *testing.T) {
	mb := Command.NewMotherBoard()
	rc := Command.NewRebotCommand(mb)
	sc := Command.NewStartCommand(mb)
	b := Command.NewBox(rc,sc)
	b.PressButton1()
	b.PressButton2()

	bb := Command.NewBatchBox()
	bb.AppendCommand(rc, sc)
	bb.Batch()
}
