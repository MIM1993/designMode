package Mediator

import "testing"

func TestMed(t *testing.T) {
	m := GetMediatorInstance()
	m.CD = &CDDriver{}
	m.CPU = &CPU{}
	m.Video = &VideoCard{}
	m.Sound = &SoundCard{}

	m.CD.InputData("music,image")
	m.CD.ReadData()

}
