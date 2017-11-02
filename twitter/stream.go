package twitter

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var finish = make(chan bool)

func getClient() *anaconda.TwitterApi {
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerKeySecret := os.Getenv("TWITTER_CONSUMER_KEY_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerKeySecret)

	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api.SetLogger(anaconda.BasicLogger)

	return api
}

func handle(stream *anaconda.Stream, c chan Tweet) {
	for {
		select {
		case item := <-stream.C:
			switch status := item.(type) {
			case anaconda.Tweet:
				t := parseTweet(status)
				c <- t
			default:
				// nop
			}
		case f := <-finish:
			if f {
				close(c)
				return
			}
		}
	}
}

func SubscribeUserStream() chan Tweet {
	c := make(chan Tweet)
	go handle(getClient().UserStream(url.Values{}), c)
	return c
}

func UnsubscribeUserStream() {
	finish <- true
}
