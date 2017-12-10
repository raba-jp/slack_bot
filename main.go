package main

import (
	"fmt"

	"github.com/raba-jp/slack_bot/slack"
	"github.com/raba-jp/slack_bot/twitter"
)

func main() {
	slackMain()
}

func slackMain() {
	api, err := slack.NewAPI()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err := api.SubscribeEventStream(); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer api.UnsubscribeEventStream()
}

func twitterMain() {
	api, err := twitter.NewAPI()
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
		case t := <-api.GetStream():
			fmt.Println(t.User.Name)
			fmt.Println(t.Text)
		default:
			// nop
		}
	}
}
