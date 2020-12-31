package Facade

import "testing"

func TestAPI(t *testing.T) {
	a := NewIAPI()
	a.Test()
}