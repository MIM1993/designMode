package Strategy

import "fmt"

type Payment struct {
	context  *PayContext
	strategy PaymentPolicy
}

type PayContext struct {
	name   string
	cardID string
	money  int64
}

func NewPayContext(name, cardID string, money int64) *PayContext {
	return &PayContext{
		name:   name,
		cardID: cardID,
		money:  money,
	}
}

func NewPayment(context *PayContext, strategy PaymentPolicy) *Payment {
	return &Payment{
		context:  context,
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	fmt.Println("开始支付....")
	p.strategy.pay(p.context.name, p.context.cardID, p.context.money)
	fmt.Println("结束支付....")
}

type PaymentPolicy interface {
	pay(name, cardID string, money int64)
}

type WXPayment struct{}

func NewWXPayment() *WXPayment {
	return &WXPayment{}
}

func (wx *WXPayment) pay(name, cardID string, money int64) {
	fmt.Printf("微信支付 -->  name:%s  cardID:%s  money:%d\n", name, cardID, money)
}

type ZFBPayment struct{}

func NewZFBPayment() *ZFBPayment {
	return &ZFBPayment{}
}

func (wx *ZFBPayment) pay(name, cardID string, money int64) {
	fmt.Printf("支付宝支付 -->  name:%s  cardID:%s  money:%d\n", name, cardID, money)
}
