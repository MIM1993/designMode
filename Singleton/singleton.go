package Singleton

import "sync"

type TmpData struct{}

var tmp *TmpData
var once = sync.Once{}

func NewTmp() *TmpData {
	once.Do(func() {
		tmp = &TmpData{}
	})
	return tmp
}
