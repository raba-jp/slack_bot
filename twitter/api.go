package twitter

import (
	"net/url"
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type Api interface {
	SubscribeUserStream() error
	UnsubscribeUserStream()
}

type ApiImpl struct {
	consumerKey       string
	consumerKeySecret string
	accessToken       string
	accessTokenSecret string
	client            *anaconda.TwitterApi
	FinishUserStream  chan bool
	Stream            chan Tweet
}

func NewApi() (*Api, error) {
	api := &Api{
		consumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		consumerKeySecret: os.Getenv("TWITTER_CONSUMER_KEY_SECRET"),
		accessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		accessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
		FinishUserStream:  make(chan bool),
		Stream:            make(chan Tweet),
	}
	if err := api.initialize(); err != nil {
		return nil, err
	}
	return api, nil
}

func (self *ApiImpl) SubscribeUserStream() error {
	if err := self.validateConfig(); err != nil {
		return err
	}
	go func() {
		stream := self.client.UserStream(url.Values{})
		for {
			select {
			case item := <-stream.C:
				switch status := item.(type) {
				case anaconda.Tweet:
					self.Stream <- parseTweet(status)
				default:
					// nop
				}
			case f := <-self.FinishUserStream:
				if f {
					return
				}
			}
		}
	}()
	return nil
}

func (self *ApiImpl) UnsubscribeUserStream() {
	self.FinishUserStream <- true
}

func (self *ApiImpl) validateConfig() error {
	if self.consumerKey == "" {
		return &TwitterConfigError{Msg: "ConsumerKey is nil"}
	}
	if self.consumerKeySecret == "" {
		return &TwitterConfigError{Msg: "ConsumerKeySecret is nil"}
	}
	if self.accessToken == "" {
		return &TwitterConfigError{Msg: "AccessToken is nil"}
	}
	if self.accessTokenSecret == "" {
		return &TwitterConfigError{Msg: "AccessTokenSecret is nil"}
	}
	return nil
}

func (self *ApiImpl) initialize() error {
	if err := self.validateConfig(); err != nil {
		return err
	}
	anaconda.SetConsumerKey(self.consumerKey)
	anaconda.SetConsumerSecret(self.consumerKeySecret)
	self.client = anaconda.NewTwitterApi(self.accessToken, self.accessTokenSecret)
	self.client.SetLogger(anaconda.BasicLogger)
	return nil
}
