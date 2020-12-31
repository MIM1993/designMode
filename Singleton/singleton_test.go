package Singleton

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTmp(t *testing.T) {
	t1 := NewTmp()
	t2 := NewTmp()
	if reflect.DeepEqual(t1,t2){
		fmt.Println("ok")
		return
	}
	fmt.Println("not ok")
}