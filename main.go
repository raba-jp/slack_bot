package main

import (
	"fmt"

	"github.com/raba-jp/twitter_client_for_slack/twitter"
)

func main() {
	stream := twitter.SubscribeUserStream()
	defer twitter.UnsubscribeUserStream()
	for {
		select {
		case <-stream:
			fmt.Println("test")
		default:
			// nop
		}
	}
}
