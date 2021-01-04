package Flyweight

import "testing"

func TestFlyweight(t *testing.T) {
	i := NewImageViewer("hello")
	ic := NewImageViewer("hello")

	if i.ImageFlyweight != ic.ImageFlyweight {
		panic(nil)
	}
	//i.Display()
}
