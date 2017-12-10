package slack

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
)

// API is Slack API interface
type API interface {
	SubscribeEventStream() error
	UnsubscribeEventStream()
}

// APIImpl is Slack API implementation
type APIImpl struct {
	apiToken     string
	client       *slack.Client
	finishStream chan bool
	stream       chan Event
}

// NewAPI is return initialized and implemented API struct
func NewAPI() (*APIImpl, error) {
	api := &APIImpl{
		apiToken:     os.Getenv("SLACK_OAUTH_ACCESS_TOKEN"),
		finishStream: make(chan bool),
		stream:       make(chan Event),
	}
	if err := api.initialize(); err != nil {
		return nil, err
	}
	return api, nil
}

// GetStream is return stream channel
func (api *APIImpl) GetStream() chan Event {
	return api.stream
}

// GetChannels is return Slack channels
func (api *APIImpl) GetChannels() ([]*Channel, error) {
	channels, err := api.client.GetChannels(false)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return nil, err
	}
	var c []*Channel
	for _, channel := range channels {
		c = append(c, &Channel{ID: channel.ID, Name: channel.Name})
	}
	return c, nil
}

// SubscribeEventStream is posted events to "EventStream" channel
func (api *APIImpl) SubscribeEventStream() error {
	rtm := api.client.NewRTM()
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

// UnsubscribeEventStream is stoped posts events to "EventStream" channel
func (api *APIImpl) UnsubscribeEventStream() {
	api.finishStream <- true
	close(api.stream)
}

// PostMessage is post message
func (api *APIImpl) PostMessage(channel Channel) {

}

func (api *APIImpl) validateConfig() error {
	if api.apiToken == "" {
		return &SlackConfigError{Msg: "API Token is nil"}
	}
	return nil
}

func (api *APIImpl) initialize() error {
	if err := api.validateConfig(); err != nil {
		return err
	}
	api.client = slack.New(api.apiToken)
	return nil
}
