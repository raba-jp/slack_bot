package main

import (
	"fmt"

	"github.com/raba-jp/twitter_client_for_slack/twitter"
)

func main() {
	api, err := twitter.NewApi()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	if err := api.SubscribeUserStream(); err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer api.UnsubscribeUserStream()
	for {
		select {
		case t := <-api.Stream:
			fmt.Println(t.User.Name)
			fmt.Println(t.Text)
		default:
			// nop
		}
	}
}
