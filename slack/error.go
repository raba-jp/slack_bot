package slack

import "fmt"

type SlackConfigError struct {
	Msg string
}

func (self *SlackConfigError) Error() string {
	return fmt.Sprintf("ERROR: %s", self.Msg)
}
