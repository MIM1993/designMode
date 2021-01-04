package Flyweight

import "fmt"

type ImageFlyweightFactory struct {
	Maps map[string]*ImageFlyweight
}

var Iff *ImageFlyweightFactory

func NewImageFlyweightFactory() *ImageFlyweightFactory {
	if Iff == nil {
		Iff = &ImageFlyweightFactory{
			Maps: make(map[string]*ImageFlyweight, 0),
		}
	}
	return Iff
}

func (i *ImageFlyweightFactory) Get(name string) *ImageFlyweight {
	if v, ok := i.Maps[name]; ok {
		return v
	}
	data := fmt.Sprintf("%s_data\n", name)
	i.Maps[name] = NewImageFlyweight(name, data)
	return i.Maps[name]
}

type ImageFlyweight struct {
	data string
	name string
}

func NewImageFlyweight(name, data string) *ImageFlyweight {
	return &ImageFlyweight{
		data: data,
		name: name,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

func (i *ImageFlyweight) Name() string {
	return i.name
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(name string) *ImageViewer {
	img := NewImageFlyweightFactory().Get(name)
	return &ImageViewer{
		ImageFlyweight: img,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("name:%s ==  data:%s\n", i.Name(), i.Data())
}
