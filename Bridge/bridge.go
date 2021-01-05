package Bridge

import "fmt"

//抽象层
type AbstractMessage interface {
	SendMessage(string, string)
}

//实现层
type MessageImplementer interface {
	Send(string, string)
}

type MessageSMS struct{}

func NewMessageSMS() MessageImplementer {
	return &MessageSMS{}
}

func (m *MessageSMS) Send(text, to string) {
	fmt.Printf("MessageSMS_text:%s to:%s\n", text, to)
}

type MessageEmail struct{}

func NewMessageEmail() MessageImplementer {
	return &MessageEmail{}
}

func (m *MessageEmail) Send(text, to string) {
	fmt.Printf("MessageEmail_text:%s to:%s\n", text, to)
}

//链接桥 实现了AbstractMessage接口,且包含了MessageImplementer接口
type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) AbstractMessage {
	return &CommonMessage{
		method: method,
	}
}

func (c *CommonMessage) SendMessage(text, to string) {
	text = "SendMessage::" + text
	c.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) AbstractMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (c *UrgencyMessage) SendMessage(text, to string) {
	text = "UrgencyMessage::" + text
	c.method.Send(text, to)
}
