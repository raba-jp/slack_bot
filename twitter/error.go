package twitter

import "fmt"

type ConfigError struct {
	Msg string
}

func (self *ConfigError) Error() string {
	return fmt.Sprintf("ERROR: %s", self.Msg)
}
