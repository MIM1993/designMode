package Proxy

//代理接口
type Subject interface {
	Do() string
}

//被代理类
type RealSubject struct {
	Data string
}

func (r *RealSubject) Do() string {
	return "real"
}

//代理类
type Proxy struct {
	real *RealSubject
}

func (p *Proxy) Do() string {
	var str string

	//调用真实类前处理流程  检查缓存，判断权限，实例化真实对象等。。
	str += "pre:"

	//调用真实类
	str += p.real.Do()

	//调用真实类后处理流程  调用之后的操作，如缓存结果，对结果进行处理等。。

	str += ":after"

	return str
}
