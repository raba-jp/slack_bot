package twitter

import "github.com/ChimeraCoder/anaconda"

func parseTweet(t anaconda.Tweet) Tweet {
	user := User{Name: t.User.Name, ScreenName: t.User.ScreenName}
	return Tweet{Text: t.ExtendedTweet.FullText, User: user}
}
