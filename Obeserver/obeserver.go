package Obeserver

import "fmt"

type Subject struct {
	obeservers []Obeserver
	context    string
}

func NewSubject(context string) *Subject {
	return &Subject{
		obeservers: make([]Obeserver, 0),
		context:    context,
	}
}

func (s *Subject) UpdateContext(str string) {
	if len(str) > 0 {
		s.context = str
	}
	s.notify()
}

func (s *Subject) notify() {
	for _, v := range s.obeservers {
		v.Update(s.context)
	}
}

func (s *Subject) AppendObeserve(obes ...Obeserver) {
	s.obeservers = append(s.obeservers, obes...)
}

//观察者调用接口
type Obeserver interface {
	Update(string)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(str string) {
	fmt.Printf("Reader:%s  context:%s\n", r.name, str)
}

func (r *Reader) Name() string {
	return r.name
}
