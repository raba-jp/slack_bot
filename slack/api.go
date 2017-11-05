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
}

func (self *Api) validateConfig() error {
	if self.apiToken == "" {
		return &SlackConfigError{Msg: "API Token is nil"}
	}
	return nil
}

func (self *Api) initialize() error {
	if err := self.validateConfig(); err != nil {
		return err
	}
	self.client = slack.New(self.apiToken)
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
				fmt.Println(ev)
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
