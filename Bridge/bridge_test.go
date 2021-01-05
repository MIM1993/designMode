package Bridge

import "testing"

func TestMessage(t *testing.T) {
	NewCommonMessage(NewMessageSMS()).SendMessage("hello","mim")
	NewCommonMessage(NewMessageEmail()).SendMessage("hello","mim")
	NewUrgencyMessage(NewMessageSMS()).SendMessage("hello","mim")
	NewUrgencyMessage(NewMessageEmail()).SendMessage("hello","mim")


}