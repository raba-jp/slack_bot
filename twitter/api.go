package twitter

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// API is Twitter User Stream API interface
type API interface {
	GetStream() chan Tweet
	SubscribeUserStream() error
	UnsubscribeUserStream()
}

// APIImpl is Twitter User Stream API implementation
type APIImpl struct {
	consumerKey       string
	consumerKeySecret string
	accessToken       string
	accessTokenSecret string
	client            *anaconda.TwitterApi
	finishStream      chan bool
	stream            chan Tweet
}

// NewAPI is return initialized and implemented API struct
func NewAPI() (*APIImpl, error) {
	api := &APIImpl{
		consumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		consumerKeySecret: os.Getenv("TWITTER_CONSUMER_KEY_SECRET"),
		accessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		accessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		finishStream:      make(chan bool),
		stream:            make(chan Tweet),
	}
	if err := api.initialize(); err != nil {
		return nil, err
	}
	return api, nil
}

// GetStream is return Stream channel
func (api *APIImpl) GetStream() chan Tweet {
	return api.stream
}

// SubscribeUserStream is posted tweet to "Stream" channel
func (api *APIImpl) SubscribeUserStream() error {
	if err := api.validateConfig(); err != nil {
		return err
	}
	go func() {
		stream := api.client.UserStream(url.Values{})
		for {
			select {
			case item := <-stream.C:
				switch status := item.(type) {
				case anaconda.Tweet:
					api.stream <- parseTweet(status)
				default:
					// nop
				}
			case f := <-api.finishStream:
				if f {
					return
				}
			}
		}
	}()
	return nil
}

// UnsubscribeUserStream is stoped posts tweet to "Stream" channel
func (api *APIImpl) UnsubscribeUserStream() {
	api.finishStream <- true
	close(api.stream)
}

func (api *APIImpl) validateConfig() error {
	if api.consumerKey == "" {
		return &TwitterConfigError{Msg: "ConsumerKey is nil"}
	}
	if api.consumerKeySecret == "" {
		return &TwitterConfigError{Msg: "ConsumerKeySecret is nil"}
	}
	if api.accessToken == "" {
		return &TwitterConfigError{Msg: "AccessToken is nil"}
	}
	if api.accessTokenSecret == "" {
		return &TwitterConfigError{Msg: "AccessTokenSecret is nil"}
	}
	return nil
}

func (api *APIImpl) initialize() error {
	if err := api.validateConfig(); err != nil {
		return err
	}
	anaconda.SetConsumerKey(api.consumerKey)
	anaconda.SetConsumerSecret(api.consumerKeySecret)
	api.client = anaconda.NewTwitterApi(api.accessToken, api.accessTokenSecret)
	api.client.SetLogger(anaconda.BasicLogger)
	return nil
}
