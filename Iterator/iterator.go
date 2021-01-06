package Iterator

//聚合类接口
type Aggregate interface {
	Iterator() Iterator
}

//针对聚合类的迭代器接口
type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

//聚合类
type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		n:    n,
		next: 0,
	}
}

//迭代器类
type NumbersIterator struct {
	n    *Numbers
	next int
}

func (n *NumbersIterator) First() {
	n.next = n.n.start
}

func (n *NumbersIterator) IsDone() bool {
	return n.next > n.n.end
}

func (n *NumbersIterator) Next() interface{} {
	if !n.IsDone() {
		next := n.next
		n.next++
		return next
	}
	return nil
}

