package twitter

import "net/url"

type User struct {
	Name       string
	ScreenName string
	IconUrl    url.URL
}
