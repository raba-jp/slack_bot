package slack

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

type Api struct {
	apiToken          string
	client            *slack.Client
	FinishEventStream chan bool
	EventStream       chan Event
	Channels          []Channel
}

func (self *Api) validateConfig() error {
	if self.apiToken == "" {
		return &SlackConfigError{Msg: "API Token is nil"}
	}
	return nil
}

func (self *Api) getChannels() []Channel {
	channels, err := self.client.GetChannels(false)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}
	var c []Channel
	for _, channel := range channels {
		append(c, &Channel{ID: channel.ID, Name: channel.Name})
	}
	return c
}

func (self *Api) initialize() error {
	if err := self.validateConfig(); err != nil {
		return err
	}
	self.client = slack.New(self.apiToken)
	self.Channels = self.getChannels()
	return nil
}

func NewApi() (*Api, error) {
	api := &Api{
		apiToken:          os.Getenv("SLACK_OAUTH_ACCESS_TOKEN"),
		FinishEventStream: make(chan bool),
		EventStream:       make(chan Event),
	}
	if err := api.initialize(); err != nil {
		return nil, err
	}
	return api, nil
}

func (self *Api) SubscribeEventStream() error {
	rtm := self.client.NewRTM()
	go rtm.ManageConnection()

	go func() {
		for msg := range rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				fmt.Printf("User: %s", ev.Channel)
			default:
				// nop
			}
		}
	}()
	return nil
}

func (self *Api) UnsubscribeEventStream() {
	self.FinishEventStream <- true
}

func (self *Api) PostMessage(channel Channel) {

}
