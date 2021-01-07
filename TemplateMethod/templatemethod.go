package TemplateMethod

import "fmt"

//抽象模板类
type DownLoader interface {
	DownLoade()
}

//具体实现类
type Implement interface {
	download()
	save()
	url() string
}

type template struct {
	Implement
}

func Newtemplate(i Implement) DownLoader {
	return &template{
		Implement: i,
	}
}

func (t *template) DownLoade() {
	fmt.Println("开始下载")
	//t.url = t.Implement.url()
	t.Implement.download()
	t.Implement.save()
	fmt.Println("结束下载")
}

func (t *template) save() {
	fmt.Println("default save model")
}

type HTTPModle struct {
	Url string
}

//
func NewHTTPModle(url string) Implement {
	h := &HTTPModle{Url: url}
	return h
}

func (h *HTTPModle) download() {
	fmt.Println("HTTPModle download modle, url:", h.Url)
}

func (h *HTTPModle) save() {
	fmt.Println("HTTPModle save modle")
}

func (h *HTTPModle) url() string {
	return h.Url
}
