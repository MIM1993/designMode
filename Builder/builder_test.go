package Builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	var b = &Builder1{}
	director := NewDirector(b)
	//建造器接口,只管构造数据,不会提供返回数据func,数据返回应由builder自行实现  ps:按个人理解
	director.Construct()
	result := b.ReturnResult()
	fmt.Println(result)
}
