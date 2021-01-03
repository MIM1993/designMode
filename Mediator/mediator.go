package Mediator

import (
	"fmt"
	"strings"
)

//初步获取数据
type CDDriver struct {
	Data string
}

func (cd *CDDriver) InputData(data string) {
	cd.Data = data
	fmt.Printf("CDDriver: input data %s\n", cd.Data)
}

func (cd *CDDriver) ReadData() {
	fmt.Printf("CDDriver: reading data %s\n", cd.Data)
	//todo 调用中介，处理数据
	GetMediatorInstance().changed(cd)
}

//cpu负责处理数据  同样调用中介
type CPU struct {
	Video string
	Sound string
}

func (cpu *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	cpu.Sound = sp[0]
	cpu.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", cpu.Sound, cpu.Video)
	//todo 调用中介，处理数据
	GetMediatorInstance().changed(cpu)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	//todo 调用中介，处理数据
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}

//中介类
type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (md *Mediator) changed(i interface{}) {
	switch t := i.(type) {
	case *CDDriver:
		md.CPU.Process(t.Data)
	case *CPU:
		md.Video.Display(t.Video)
		md.Sound.Play(t.Sound)
	case *VideoCard:
		fmt.Println("==Video==")
	case *SoundCard:
		fmt.Println("==Sound==")
	}
}
