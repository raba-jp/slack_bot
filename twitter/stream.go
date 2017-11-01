package twitter

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

var finish = make(chan bool)

func getClient() anaconda.TwitterApi {
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
	select {
	case item := steam.C:
		switch status := item.(type) {
		case anaconda.Tweet:
			t := parseTweet(status)
			c <- t
		default:
			// nop
		}
	}
}

func SubscribeUserStream() chan Tweet {
	c := make(chan Tweet)
	go func() {
		for {
			handle(getClient().UserStream(url.Values{}), c)

			select {
			case f := <-finish:
				if f {
					close(c)
					return
				}
			default:
				// nop
			}
		}
	}()
	return c
}

func UnsubscribeUserStream() {
	finish <- true
}

func main() {
}
