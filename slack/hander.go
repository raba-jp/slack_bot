package slack

import "github.com/nlopes/slack"

type MessageEventHandler struct {
	channels []string
}

func (self *MessageEventHandler) validateMention(ev slack.Event) error {
	return nil
}

func (self *MessageEventHandler) validateChannel(ev slack.Event) error {
	return nil
}
