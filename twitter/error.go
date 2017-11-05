package twitter

import "fmt"

type TwitterConfigError struct {
	Msg string
}

func (self *TwitterConfigError) Error() string {
	return fmt.Sprintf("ERROR: %s", self.Msg)
}
