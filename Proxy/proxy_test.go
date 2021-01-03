package Proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}
	fmt.Println(sub.Do())
}
