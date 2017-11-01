package twitter

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

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

func ListenUserStream() chan Tweet {
	c := make(chan Tweet)
	steam := getClient().UserStream(url.Values{})
	go func() {
		for {
			select {
			case item := steam.C:
				switch status := item.(type) {
				case anaconda.Tweet:
					parseTweet(status)
				}
			}
		}
	}()
}

func main() {
}
