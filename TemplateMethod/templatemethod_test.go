package TemplateMethod

import "testing"

func TestNewHTTPModle(t *testing.T) {
	d := NewHTTPModle("http://example.com/abc.zip")

	var tm DownLoader
	tm = Newtemplate(d)
	tm.DownLoade()
	//d.DownLoade("http://example.com/abc.zip")
}